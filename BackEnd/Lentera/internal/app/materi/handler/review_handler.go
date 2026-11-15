package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

type ReviewHandler struct {
	svc contract.ReviewService
}

func NewReview(svc contract.ReviewService) *ReviewHandler {
	return &ReviewHandler{svc: svc}
}

// GET /cohorts/:id/reviews?status=&material_id=&q=&include_inactive=
func (h *ReviewHandler) List(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}

	var q dto.ListReviewsQuery
	if err := c.ShouldBindQuery(&q); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.ListReviews(c.Request.Context(), userID, cohortID, q)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// GET /cohorts/:id/progress/:progressId
func (h *ReviewHandler) Detail(c *gin.Context) {
	userID, cohortID, progressID, ok := reviewParams(c)
	if !ok {
		return
	}

	res, err := h.svc.GetReviewItem(c.Request.Context(), userID, cohortID, progressID)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// POST /cohorts/:id/progress/:progressId/approve
func (h *ReviewHandler) Approve(c *gin.Context) {
	userID, cohortID, progressID, ok := reviewParams(c)
	if !ok {
		return
	}

	// body opsional: kosong boleh, berarti approve tanpa catatan
	var req dto.ApproveRequest
	if c.Request.ContentLength > 0 {
		if err := c.ShouldBindJSON(&req); err != nil {
			writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
			return
		}
	}

	res, err := h.svc.Approve(c.Request.Context(), userID, cohortID, progressID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// POST /cohorts/:id/progress/:progressId/request-revision
func (h *ReviewHandler) RequestRevision(c *gin.Context) {
	userID, cohortID, progressID, ok := reviewParams(c)
	if !ok {
		return
	}

	var req dto.ReviewNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.RequestRevision(c.Request.Context(), userID, cohortID, progressID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// POST /cohorts/:id/progress/:progressId/reopen
func (h *ReviewHandler) Reopen(c *gin.Context) {
	userID, cohortID, progressID, ok := reviewParams(c)
	if !ok {
		return
	}

	var req dto.ReviewNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		writeError(c, fmt.Errorf("%w: %v", response.ErrBadRequest, err))
		return
	}

	res, err := h.svc.Reopen(c.Request.Context(), userID, cohortID, progressID, req)
	if err != nil {
		writeError(c, err)
		return
	}
	writeOK(c, res)
}

// ambil user, cohort id, progress id sekaligus; sudah menulis error kalau gagal
func reviewParams(c *gin.Context) (userID, cohortID, progressID uint, ok bool) {
	userID, ok = currentUser(c)
	if !ok {
		return 0, 0, 0, false
	}
	var err error
	if cohortID, err = paramID(c, "id"); err != nil {
		writeError(c, err)
		return 0, 0, 0, false
	}
	if progressID, err = paramID(c, "progressId"); err != nil {
		writeError(c, err)
		return 0, 0, 0, false
	}
	return userID, cohortID, progressID, true
}
