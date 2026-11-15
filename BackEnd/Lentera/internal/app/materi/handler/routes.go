package handler

import "github.com/gin-gonic/gin"

// rg = group yang SUDAH dipasang middleware Auth (JWT)
func (h *MaterialHandler) RegisterRoutes(rg *gin.RouterGroup) {
	cohort := rg.Group("/cohorts/:cohortId")

	// materi & progres pribadi (Member & Mentor)
	cohort.GET("/materials", h.List)
	cohort.GET("/materials/:materialId", h.Detail)
	cohort.GET("/me/summary", h.MySummary)

	// nanti: CRUD materi (Mentor), start/submit (Member), review (Mentor)
}
