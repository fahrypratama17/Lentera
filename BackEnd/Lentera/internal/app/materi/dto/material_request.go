package dto

import (
	"encoding/json"
	"time"
)

type ListMaterialsQuery struct {
	Q      string `form:"q" binding:"omitempty,max=100"`
	Status string `form:"status" binding:"omitempty,oneof=not_started in_progress submitted completed"`
	Late   *bool  `form:"late"`
	PaginationQuery
}

// POST /cohorts/:cohortId/materials
type CreateMaterialRequest struct {
	Title      string     `json:"title" binding:"required,min=3,max=200"`
	Guide      string     `json:"guide" binding:"required,max=10000"`
	References []string   `json:"references" binding:"omitempty,max=20,dive,url"`
	Deadline   *time.Time `json:"deadline"`
}

// PATCH /cohorts/:cohortId/materials/:materialId — semua field opsional
type UpdateMaterialRequest struct {
	Title      *string      `json:"title" binding:"omitempty,min=3,max=200"`
	Guide      *string      `json:"guide" binding:"omitempty,max=10000"`
	References []string     `json:"references" binding:"omitempty,max=20,dive,url"`
	Deadline   NullableTime `json:"deadline"`
}

// PUT /cohorts/:cohortId/materials/order
type ReorderMaterialsRequest struct {
	MaterialIDs []uint `json:"material_ids" binding:"required,min=1,dive,gt=0"`
}

// NullableTime membedakan 3 kondisi field JSON:
//   - tidak dikirim         → Set=false              (jangan ubah)
//   - "deadline": null      → Set=true, Value=nil    (hapus deadline)
//   - "deadline": "2026-…"  → Set=true, Value=&t     (ganti deadline)
type NullableTime struct {
	Set   bool
	Value *time.Time
}

func (n *NullableTime) UnmarshalJSON(b []byte) error {
	n.Set = true // fungsi ini hanya dipanggil kalau key "deadline" ada di JSON
	if string(b) == "null" {
		n.Value = nil
		return nil
	}
	var t time.Time
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}
	n.Value = &t
	return nil
}
