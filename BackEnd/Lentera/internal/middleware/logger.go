package middleware

import (
	"crypto/rand"
	"encoding/hex"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger: access log JSON per request, dengan request_id untuk menelusuri request.
func Logger(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqID := c.GetHeader("X-Request-ID")
		if reqID == "" {
			reqID = newRequestID()
		}
		c.Set("request_id", reqID)
		c.Header("X-Request-ID", reqID) // dikirim balik ke client, berguna saat debug bareng FE

		c.Next()

		if c.Request.URL.Path == "/metrics" {
			return // jangan penuhi log dengan scrape Prometheus
		}

		route := c.FullPath()
		if route == "" {
			route = "unmatched"
		}
		status := c.Writer.Status()

		attrs := []any{
			"request_id", reqID,
			"method", c.Request.Method,
			"route", route,
			"status", status,
			"latency_ms", time.Since(start).Milliseconds(),
			"client_ip", c.ClientIP(),
		}
		if uid := c.GetUint("uid"); uid != 0 {
			attrs = append(attrs, "user_id", uid)
		}
		if len(c.Errors) > 0 {
			attrs = append(attrs, "errors", c.Errors.String())
		}

		switch {
		case status >= 500:
			log.Error("request", attrs...)
		case status >= 400:
			log.Warn("request", attrs...)
		default:
			log.Info("request", attrs...)
		}
	}
}

func newRequestID() string {
	b := make([]byte, 8)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
