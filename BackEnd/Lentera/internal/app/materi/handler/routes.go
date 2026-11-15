package handler

import "github.com/gin-gonic/gin"

func (h *MaterialHandler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	cohort := api.Group("/cohorts/:id", auth)

	// materi & progres pribadi (Member & Mentor)
	cohort.GET("/materials", h.List)
	cohort.GET("/materials/:materialId", h.Detail)
	cohort.GET("/me/summary", h.MySummary)

	// CRUD materi (Mentor — dicek di service)
	cohort.POST("/materials", h.Create)
	cohort.PUT("/materials/order", h.Reorder)
	cohort.PATCH("/materials/:materialId", h.Update)
	cohort.DELETE("/materials/:materialId", h.Delete)

	// progres & pengumpulan (Member)
	cohort.POST("/materials/:materialId/start", h.Start)
	cohort.POST("/materials/:materialId/submit", h.Submit)

	// review pengumpulan (Mentor)
	cohort.POST("/materials/:materialId/submissions/:userId/review", h.Review)

	// upload file PDF ke Supabase (Mentor: purpose=material, Member: purpose=submission)
	cohort.POST("/materials/upload", h.UploadPDF)
}

func (h *ReviewHandler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	cohort := api.Group("/cohorts/:id", auth)

	cohort.GET("/reviews", h.List)
	cohort.GET("/progress/:progressId", h.Detail)
	cohort.POST("/progress/:progressId/approve", h.Approve)
	cohort.POST("/progress/:progressId/request-revision", h.RequestRevision)
	cohort.POST("/progress/:progressId/reopen", h.Reopen)
}

func (h *MonitoringHandler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	cohort := api.Group("/cohorts/:id", auth)
	cohort.GET("/dashboard", h.Dashboard)
	cohort.GET("/matrix", h.Matrix)
	cohort.GET("/members/:userId/progress", h.MemberProgress)
}
