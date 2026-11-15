package dto

// GET /cohorts/:cohortId/matrix
type MatrixQuery struct {
	Q          string `form:"q" binding:"omitempty,max=100"`                                                // cari nama member
	MaterialID uint   `form:"material_id" binding:"omitempty,gt=0"`                                         // tampilkan 1 kolom materi saja
	Status     string `form:"status" binding:"omitempty,oneof=not_started in_progress submitted completed"` // member yang punya ≥1 sel dengan status ini
}
