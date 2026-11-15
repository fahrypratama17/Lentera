package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/cohort/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type Handler struct {
	svc contract.CohortService
}

func New(svc contract.CohortService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) CreateCohort(c *gin.Context) {
	var req dto.CreateCohortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	userID := c.GetUint("user_id")
	res, err := h.svc.CreateCohort(c.Request.Context(), userID, req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusCreated, res)
}

func (h *Handler) ListMyCohorts(c *gin.Context) {
	userID := c.GetUint("user_id")
	res, err := h.svc.ListMyCohorts(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) InviteMember(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	var req dto.CreateInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	userID := c.GetUint("user_id")
	res, err := h.svc.InviteMember(c.Request.Context(), userID, uint(cohortID), req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusCreated, res)
}

func (h *Handler) ListCohortInvitations(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	userID := c.GetUint("user_id")
	res, err := h.svc.ListCohortInvitations(c.Request.Context(), userID, uint(cohortID))
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) CancelInvitation(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	invID, err := strconv.ParseUint(c.Param("invitationId"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	userID := c.GetUint("user_id")
	if err := h.svc.CancelInvitation(c.Request.Context(), userID, uint(cohortID), uint(invID)); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "invitation cancelled"})
}

func (h *Handler) ListMyInvitations(c *gin.Context) {
	userID := c.GetUint("user_id")
	res, err := h.svc.ListMyInvitations(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) RespondInvitation(c *gin.Context) {
	invID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	var req dto.RespondInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	userID := c.GetUint("user_id")
	if err := h.svc.RespondInvitation(c.Request.Context(), userID, uint(invID), req); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "invitation responded"})
}

func (h *Handler) GetCohortDetail(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	userID := c.GetUint("user_id")
	res, err := h.svc.GetCohortDetail(c.Request.Context(), userID, uint(cohortID))
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) UpdateCohort(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	var req dto.UpdateCohortRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	userID := c.GetUint("user_id")
	res, err := h.svc.UpdateCohort(c.Request.Context(), userID, uint(cohortID), req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) DeleteCohort(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	userID := c.GetUint("user_id")
	if err := h.svc.DeleteCohort(c.Request.Context(), userID, uint(cohortID)); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "cohort deleted"})
}

func (h *Handler) ListMembers(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	userID := c.GetUint("user_id")
	res, err := h.svc.ListMembers(c.Request.Context(), userID, uint(cohortID))
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) UpdateMemberRole(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	targetUserID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	var req dto.UpdateMemberRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	actorID := c.GetUint("user_id")
	if err := h.svc.UpdateMemberRole(c.Request.Context(), actorID, uint(cohortID), uint(targetUserID), req); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "member role updated"})
}

func (h *Handler) LeaveCohort(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	userID := c.GetUint("user_id")
	if err := h.svc.LeaveCohort(c.Request.Context(), userID, uint(cohortID)); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "left cohort"})
}

func (h *Handler) RemoveMember(c *gin.Context) {
	cohortID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	targetUserID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		response.Error(c, response.ErrBadRequest)
		return
	}
	actorID := c.GetUint("user_id")
	if err := h.svc.RemoveMember(c.Request.Context(), actorID, uint(cohortID), uint(targetUserID)); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "member removed"})
}