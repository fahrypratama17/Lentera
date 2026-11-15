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

}

func (h *MonitoringHandler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	cohort := api.Group("/cohorts/:id", auth)

	// pemantauan mentor (role dicek di service)
	cohort.GET("/dashboard", h.Dashboard)
	cohort.GET("/matrix", h.Matrix)
	cohort.GET("/members/:userId/progress", h.MemberProgress)
}
