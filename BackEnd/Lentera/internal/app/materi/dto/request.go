package dto

type ListMaterialsQuery struct {
	Q      string `form:"q" binding:"omitempty,max=100"`
	Status string `form:"status" binding:"omitempty,oneof=not_started in_progress submitted completed"`
	Late   *bool  `form:"late"`
}
