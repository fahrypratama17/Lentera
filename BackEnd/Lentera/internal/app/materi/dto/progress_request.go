package dto

type SubmitMaterialRequest struct {
	Content string  `json:"content" binding:"omitempty,max=20000"`
	WorkURL *string `json:"work_url" binding:"omitempty,url,max=500"`
}

type ReviewSubmissionRequest struct {
	Decision string  `json:"decision" binding:"required,oneof=approved revision_requested"`
	Note     *string `json:"note" binding:"omitempty,max=5000"`
}