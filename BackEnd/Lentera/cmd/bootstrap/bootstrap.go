package bootstrap

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/config"
	cohortHandler "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/handler"
	cohortRepo "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/repository"
	cohortService "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/service"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	materialHandler "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/handler"
	materialRepo "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/repository"
	materialService "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/service"
	userHandler "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/handler"
	userRepo "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/repository"
	userService "github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/service"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/infra/cache"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/infra/database"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/middleware"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/realtime"

	"github.com/gin-gonic/gin"
)

func Run() {
	cfg := config.Load()
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// ---- infra ----
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Error("db connect failed", "error", err)
		os.Exit(1)
	}
	log.Info("db connected")

	if err := database.Migrate(db); err != nil {
		log.Error("migrate failed", "error", err)
		os.Exit(1)
	}
	if err := database.Seed(db); err != nil {
		log.Error("seed failed", "error", err)
	}

	rdb, err := cache.Connect(cfg.RedisAddr)
	if err != nil {
		log.Warn("redis unavailable, cache disabled", "error", err)
		rdb = nil
	}
	_ = rdb // TODO: dipakai untuk cache materi/dashboard
	hub := realtime.NewHub()

	// ---- http ----
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logger(log), middleware.Metrics(), middleware.CORS(cfg.CORSOrigin))

	r.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.PingContext(c.Request.Context()) != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "db down"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	// TODO: r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// TODO: r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api/v1")
	auth := middleware.Auth(cfg.JWTSecret)

	// ---- modules (DI manual: repo -> service -> handler) ----
	userHandler.New(userService.New(userRepo.New(db), cfg.JWTSecret)).RegisterRoutes(api, auth)
	cohortHandler.New(cohortService.New(cohortRepo.New(db))).RegisterRoutes(api, auth)

	now := func() time.Time { return time.Now().UTC() }
	materialHandler.New(materialService.New(materialRepo.New(db), contract.NoopBroadcaster{}, now)).RegisterRoutes(api, auth)
	materialHandler.NewMonitoring(materialService.NewMonitoring(materialRepo.NewMonitoring(db), now)).RegisterRoutes(api, auth)
	materialHandler.NewReview(materialService.NewReview(materialRepo.NewReview(db), contract.NoopBroadcaster{}, now)).RegisterRoutes(api, auth)
	// TODO: modul baru -> tambah 1 baris di sini

	api.GET("/ws", auth, hub.ServeWS)

	// ---- server + graceful shutdown ----
	srv := &http.Server{Addr: ":" + cfg.Port, Handler: r, ReadHeaderTimeout: 10 * time.Second}
	go func() {
		log.Info("server started", "port", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server error", "error", err)
			os.Exit(1)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Info("server stopped")
}
