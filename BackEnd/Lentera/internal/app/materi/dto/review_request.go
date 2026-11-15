package dto

// GET /cohorts/:id/reviews
type ListReviewsQuery struct {
	Status          string `form:"status" binding:"omitempty,oneof=submitted in_progress completed"` // default: submitted
	MaterialID      uint   `form:"material_id" binding:"omitempty,gt=0"`
	Q               string `form:"q" binding:"omitempty,max=100"` // cari nama member
	IncludeInactive bool   `form:"include_inactive"`              // sertakan hasil mantan anggota
}

// POST .../approve — catatan opsional
type ApproveRequest struct {
	Note *string `json:"note" binding:"omitempty,max=2000"`
}

// POST .../request-revision dan .../reopen — catatan wajib
type ReviewNoteRequest struct {
	Note string `json:"note" binding:"required,min=1,max=2000"`
}
