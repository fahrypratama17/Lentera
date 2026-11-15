package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type MonitoringHandler struct {
	svc contract.MonitoringService
}

func NewMonitoring(svc contract.MonitoringService) *MonitoringHandler {
	return &MonitoringHandler{svc: svc}
}

// GET /cohorts/:cohortId/dashboard
func (h *MonitoringHandler) Dashboard(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "cohortId")
	if err != nil {
		writeError(c, err)
		return
	}

	res, err := h.svc.Dashboard(c.Request.Context(), userID, cohortID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// GET /cohorts/:cohortId/matrix?q=&material_id=&status=
func (h *MonitoringHandler) Matrix(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "cohortId")
	if err != nil {
		writeError(c, err)
		return
	}

	var q dto.MatrixQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.Matrix(c.Request.Context(), userID, cohortID, q)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// GET /cohorts/:cohortId/members/:userId/progress
func (h *MonitoringHandler) MemberProgress(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "cohortId")
	if err != nil {
		writeError(c, err)
		return
	}
	memberID, err := paramID(c, "userId")
	if err != nil {
		writeError(c, err)
		return
	}

	res, err := h.svc.MemberProgress(c.Request.Context(), userID, cohortID, memberID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}
