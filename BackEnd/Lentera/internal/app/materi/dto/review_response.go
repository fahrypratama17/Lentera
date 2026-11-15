package dto

// ProgressResponse di-embed → field-nya (id, status, submission, last_review, ...)
// muncul langsung di level atas JSON, ditambah member & material.
type ReviewItemResponse struct {
	ProgressResponse
	Member   ReviewMember   `json:"member"`
	Material MatrixMaterial `json:"material"`
}

type ReviewMember struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Status string `json:"status"` // active | left
}
