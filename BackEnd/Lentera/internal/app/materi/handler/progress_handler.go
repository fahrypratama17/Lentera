package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// POST /cohorts/:id/materials/:materialId/start
func (h *MaterialHandler) Start(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}
	materialID, err := paramID(c, "materialId")
	if err != nil {
		writeError(c, err)
		return
	}

	res, err := h.svc.StartMaterial(c.Request.Context(), userID, cohortID, materialID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// POST /cohorts/:id/materials/:materialId/submit
func (h *MaterialHandler) Submit(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}
	materialID, err := paramID(c, "materialId")
	if err != nil {
		writeError(c, err)
		return
	}

	var req dto.SubmitMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.SubmitMaterial(c.Request.Context(), userID, cohortID, materialID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// POST /cohorts/:id/materials/:materialId/submissions/:userId/review
func (h *MaterialHandler) Review(c *gin.Context) {
	mentorID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}
	materialID, err := paramID(c, "materialId")
	if err != nil {
		writeError(c, err)
		return
	}
	targetUserID, err := paramID(c, "userId")
	if err != nil {
		writeError(c, err)
		return
	}

	var req dto.ReviewSubmissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.ReviewSubmission(c.Request.Context(), mentorID, cohortID, materialID, targetUserID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}