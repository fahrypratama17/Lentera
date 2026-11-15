package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type MaterialHandler struct {
	svc contract.MaterialService
}

func NewMaterialHandler(svc contract.MaterialService) *MaterialHandler {
	return &MaterialHandler{svc: svc}
}

// GET /cohorts/:cohortId/materials?q=&status=&late=
func (h *MaterialHandler) List(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "cohortId")
	if err != nil {
		writeError(c, err)
		return
	}

	var q dto.ListMaterialsQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.ListMaterials(c.Request.Context(), userID, cohortID, q)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// GET /cohorts/:cohortId/materials/:materialId
func (h *MaterialHandler) Detail(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "cohortId")
	if err != nil {
		writeError(c, err)
		return
	}
	materialID, err := paramID(c, "materialId")
	if err != nil {
		writeError(c, err)
		return
	}

	res, err := h.svc.GetMaterial(c.Request.Context(), userID, cohortID, materialID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// GET /cohorts/:cohortId/me/summary
func (h *MaterialHandler) MySummary(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "cohortId")
	if err != nil {
		writeError(c, err)
		return
	}

	res, err := h.svc.MySummary(c.Request.Context(), userID, cohortID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}
