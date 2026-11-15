package dto

import "time"

// ───────────── dashboard ─────────────

type DashboardResponse struct {
	ActiveMemberCount  int     `json:"active_member_count"`
	MaterialCount      int     `json:"material_count"`
	PendingReviewCount int     `json:"pending_review_count"`
	CompletedCount     int     `json:"completed_count"`
	CompletionRate     float64 `json:"completion_rate"`
	LateCount          int     `json:"late_count"`
}

// ───────────── matriks ─────────────

type MatrixResponse struct {
	Materials []MatrixMaterial `json:"materials"`
	Rows      []MatrixRow      `json:"rows"`
	Meta      PaginationMeta   `json:"meta"`
}

type MatrixMaterial struct {
	ID       uint       `json:"id"`
	Title    string     `json:"title"`
	Position int        `json:"position"`
	Deadline *time.Time `json:"deadline"`
}

type MatrixRow struct {
	UserID         uint         `json:"user_id"`
	Name           string       `json:"name"`
	Completed      int          `json:"completed"`
	CompletionRate float64      `json:"completion_rate"`
	Cells          []MatrixCell `json:"cells"` // urutan sama dengan Materials
}

type MatrixCell struct {
	MaterialID uint   `json:"material_id"`
	ProgressID *uint  `json:"progress_id"` // null = belum pernah dimulai
	Status     string `json:"status"`
	IsLate     bool   `json:"is_late"`
}

// ───────────── hasil belajar per member ─────────────

type MemberProgressResponse struct {
	Member  MemberBrief          `json:"member"`
	Summary PersonalSummary      `json:"summary"`
	Items   []MemberProgressItem `json:"items"`
}

type MemberBrief struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Status string `json:"status"` // active | left
}

type MemberProgressItem struct {
	Material MatrixMaterial   `json:"material"`
	Progress ProgressResponse `json:"progress"` // berisi submission & last_review
}
