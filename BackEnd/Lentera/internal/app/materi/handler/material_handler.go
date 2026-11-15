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

func New(svc contract.MaterialService) *MaterialHandler {
	return &MaterialHandler{svc: svc}
}

// GET /cohorts/:id/materials?q=&status=&late=
func (h *MaterialHandler) List(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
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

// GET /cohorts/:id/materials/:materialId
func (h *MaterialHandler) Detail(c *gin.Context) {
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

	res, err := h.svc.GetMaterial(c.Request.Context(), userID, cohortID, materialID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// GET /cohorts/:id/me/summary
func (h *MaterialHandler) MySummary(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
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

// ───────────────── CRUD materi (Mentor) ─────────────────

// POST /cohorts/:id/materials
func (h *MaterialHandler) Create(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}

	var req dto.CreateMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.CreateMaterial(c.Request.Context(), userID, cohortID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeCreated(c, res)
}

// PATCH /cohorts/:id/materials/:materialId
func (h *MaterialHandler) Update(c *gin.Context) {
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

	var req dto.UpdateMaterialRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.UpdateMaterial(c.Request.Context(), userID, cohortID, materialID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// DELETE /cohorts/:id/materials/:materialId
func (h *MaterialHandler) Delete(c *gin.Context) {
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

	if err := h.svc.DeleteMaterial(c.Request.Context(), userID, cohortID, materialID); err != nil {
		writeError(c, err)
		return
	}
	writeNoContent(c)
}

// PUT /cohorts/:id/materials/order
func (h *MaterialHandler) Reorder(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}

	var req dto.ReorderMaterialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.ReorderMaterials(c.Request.Context(), userID, cohortID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}
