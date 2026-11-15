package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	cohorts := api.Group("/cohorts", auth)
	cohorts.GET("", h.ListMyCohorts)
	cohorts.POST("", h.CreateCohort)
	cohorts.GET("/:id", h.GetCohortDetail)
	cohorts.PUT("/:id", h.UpdateCohort)
	cohorts.DELETE("/:id", h.DeleteCohort)

	// Anggota & Pengaturan Keanggotaan
	cohorts.GET("/:id/members", h.ListMembers)
	cohorts.PATCH("/:id/members/:userId/role", h.UpdateMemberRole)
	cohorts.POST("/:id/leave", h.LeaveCohort)
	cohorts.DELETE("/:id/members/:userId", h.RemoveMember)

	// Undangan (dari sisi Mentor)
	cohorts.GET("/:id/invitations", h.ListCohortInvitations)
	cohorts.POST("/:id/invitations", h.InviteMember)
	cohorts.DELETE("/:id/invitations/:invitationId", h.CancelInvitation)

	// Undangan (dari sisi Member)
	invitations := api.Group("/invitations", auth)
	invitations.GET("", h.ListMyInvitations)
	invitations.POST("/:id/respond", h.RespondInvitation)
}