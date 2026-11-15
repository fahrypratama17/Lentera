package service

import (
	"context"
	"errors"
	"testing"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── StartMaterial ─────────────────

func TestStartMaterial(t *testing.T) {
	t.Run("member mulai materi baru", func(t *testing.T) {
		svc, _, bc := newCRUDService()

		got, err := svc.StartMaterial(context.Background(), 3, 1, 11)
		if err != nil {
			t.Fatal(err)
		}
		if got.Status != entity.StatusInProgress {
			t.Fatalf("want status in_progress, got %s", got.Status)
		}
		if got.StartedAt == nil || !got.StartedAt.Equal(fixedNow) {
			t.Errorf("want started_at = fixedNow, got %v", got.StartedAt)
		}
		if !bc.has("cohort:1", "progress.updated") {
			t.Error("harus broadcast progress.updated ke room cohort")
		}
		if !bc.has("cohort:1:mentors", "dashboard.updated") {
			t.Error("harus broadcast dashboard.updated ke room mentor")
		}
	})

	t.Run("idempoten jika sudah in_progress", func(t *testing.T) {
		svc, _, bc := newCRUDService()

		got, err := svc.StartMaterial(context.Background(), 2, 1, 11)
		if err != nil {
			t.Fatal(err)
		}
		if got.Status != entity.StatusInProgress {
			t.Fatalf("want in_progress, got %s", got.Status)
		}
		if len(bc.sent) != 0 {
			t.Errorf("tidak perlu broadcast ulang jika sudah in_progress, got %v", bc.sent)
		}
	})

	t.Run("ditolak jika sudah submitted atau mentor", func(t *testing.T) {
		tests := []struct {
			name       string
			userID     uint
			materialID uint
			wantErr    error
		}{
			{name: "sudah submitted", userID: 2, materialID: 10, wantErr: response.ErrConflict},
			{name: "sudah completed", userID: 2, materialID: 12, wantErr: response.ErrConflict},
			{name: "mentor tidak bisa start", userID: 1, materialID: 11, wantErr: response.ErrForbidden},
			{name: "materi cohort lain", userID: 3, materialID: 20, wantErr: response.ErrNotFound},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				svc, _, _ := newCRUDService()
				_, err := svc.StartMaterial(context.Background(), tt.userID, 1, tt.materialID)
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("want %v, got %v", tt.wantErr, err)
				}
			})
		}
	})
}

// ───────────────── SubmitMaterial ─────────────────

func TestSubmitMaterial(t *testing.T) {
	t.Run("member submit uraian dan work_url", func(t *testing.T) {
		svc, _, bc := newCRUDService()

		got, err := svc.SubmitMaterial(context.Background(), 2, 1, 11, dto.SubmitMaterialRequest{
			Content: "  Sudah selesai refactoring GORM  ",
			WorkURL: strPtr("https://github.com/lentera/pr-11"),
		})
		if err != nil {
			t.Fatal(err)
		}
		if got.Status != entity.StatusSubmitted {
			t.Fatalf("want submitted, got %s", got.Status)
		}
		if got.Submission == nil || got.Submission.Content != "Sudah selesai refactoring GORM" {
			t.Fatalf("submission tidak tersimpan dengan benar: %+v", got.Submission)
		}
		if got.Submission.WorkURL == nil || *got.Submission.WorkURL != "https://github.com/lentera/pr-11" {
			t.Errorf("work_url tidak sesuai: %v", got.Submission.WorkURL)
		}
		if !bc.has("cohort:1", "progress.updated") || !bc.has("cohort:1:mentors", "dashboard.updated") {
			t.Error("harus broadcast progress.updated dan dashboard.updated")
		}
	})

	t.Run("member submit hanya work_url tanpa content", func(t *testing.T) {
		svc, _, _ := newCRUDService()

		got, err := svc.SubmitMaterial(context.Background(), 2, 1, 11, dto.SubmitMaterialRequest{
			WorkURL: strPtr("https://github.com/lentera/pr-11"),
		})
		if err != nil {
			t.Fatal(err)
		}
		if got.Submission == nil || got.Submission.WorkURL == nil {
			t.Fatalf("submission harus tetap terisi meski content kosong: %+v", got.Submission)
		}
	})

	t.Run("ditolak pada kondisi tidak valid", func(t *testing.T) {
		tests := []struct {
			name       string
			userID     uint
			materialID uint
			req        dto.SubmitMaterialRequest
			wantErr    error
		}{
			{
				name:       "content dan work_url kosong",
				userID:     2,
				materialID: 11,
				req:        dto.SubmitMaterialRequest{Content: "   "},
				wantErr:    response.ErrBadRequest,
			},
			{
				name:       "belum mulai materi",
				userID:     3,
				materialID: 11,
				req:        dto.SubmitMaterialRequest{Content: "langsung kumpul"},
				wantErr:    response.ErrConflict,
			},
			{
				name:       "sedang menunggu review (submitted)",
				userID:     2,
				materialID: 10,
				req:        dto.SubmitMaterialRequest{Content: "kirim ulang"},
				wantErr:    response.ErrConflict,
			},
			{
				name:       "mentor tidak bisa submit",
				userID:     1,
				materialID: 11,
				req:        dto.SubmitMaterialRequest{Content: "tes"},
				wantErr:    response.ErrForbidden,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				svc, _, bc := newCRUDService()
				_, err := svc.SubmitMaterial(context.Background(), tt.userID, 1, tt.materialID, tt.req)
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("want %v, got %v", tt.wantErr, err)
				}
				if len(bc.sent) != 0 {
					t.Errorf("tidak boleh broadcast saat gagal, got %v", bc.sent)
				}
			})
		}
	})
}

// ───────────────── ReviewSubmission ─────────────────

func TestReviewSubmission(t *testing.T) {
	t.Run("mentor menyetujui submission", func(t *testing.T) {
		svc, _, bc := newCRUDService()

		got, err := svc.ReviewSubmission(context.Background(), 1, 1, 10, 2, dto.ReviewSubmissionRequest{
			Decision: DecisionApproved,
			Note:     strPtr("  Bagus, lanjut materi berikutnya  "),
		})
		if err != nil {
			t.Fatal(err)
		}
		if got.Status != entity.StatusCompleted {
			t.Fatalf("want completed, got %s", got.Status)
		}
		if got.CompletedAt == nil || !got.CompletedAt.Equal(fixedNow) {
			t.Errorf("want completed_at = fixedNow, got %v", got.CompletedAt)
		}
		if got.LastReview == nil || got.LastReview.Decision != DecisionApproved {
			t.Fatalf("want last_review approved, got %+v", got.LastReview)
		}
		if got.LastReview.Reviewer.Name != "Mentor Satu" {
			t.Errorf("want nama reviewer Mentor Satu, got %q", got.LastReview.Reviewer.Name)
		}
		if !bc.has("cohort:1", "progress.updated") || !bc.has("cohort:1:mentors", "dashboard.updated") {
			t.Error("harus broadcast progress.updated dan dashboard.updated")
		}
	})

	t.Run("mentor meminta revisi mengembalikan status ke in_progress", func(t *testing.T) {
		svc, _, _ := newCRUDService()

		got, err := svc.ReviewSubmission(context.Background(), 1, 1, 10, 2, dto.ReviewSubmissionRequest{
			Decision: DecisionRevisionRequested,
			Note:     strPtr("Tambahkan validasi input"),
		})
		if err != nil {
			t.Fatal(err)
		}
		if got.Status != entity.StatusInProgress {
			t.Fatalf("want in_progress, got %s", got.Status)
		}
		if got.CompletedAt != nil {
			t.Errorf("completed_at harus nil saat revisi, got %v", got.CompletedAt)
		}
	})

	t.Run("review ditolak", func(t *testing.T) {
		tests := []struct {
			name         string
			mentorID     uint
			materialID   uint
			targetUserID uint
			req          dto.ReviewSubmissionRequest
			wantErr      error
		}{
			{
				name:         "revisi tanpa catatan",
				mentorID:     1,
				materialID:   10,
				targetUserID: 2,
				req:          dto.ReviewSubmissionRequest{Decision: DecisionRevisionRequested},
				wantErr:      response.ErrBadRequest,
			},
			{
				name:         "approve materi yang belum submitted",
				mentorID:     1,
				materialID:   11,
				targetUserID: 2,
				req:          dto.ReviewSubmissionRequest{Decision: DecisionApproved},
				wantErr:      response.ErrConflict,
			},
			{
				name:         "member mencoba mereview",
				mentorID:     3,
				materialID:   10,
				targetUserID: 2,
				req:          dto.ReviewSubmissionRequest{Decision: DecisionApproved},
				wantErr:      response.ErrForbidden,
			},
			{
				name:         "target adalah mentor",
				mentorID:     1,
				materialID:   10,
				targetUserID: 1,
				req:          dto.ReviewSubmissionRequest{Decision: DecisionApproved},
				wantErr:      response.ErrBadRequest,
			},
			{
				name:         "member belum memulai materi",
				mentorID:     1,
				materialID:   11,
				targetUserID: 3,
				req:          dto.ReviewSubmissionRequest{Decision: DecisionApproved},
				wantErr:      response.ErrNotFound,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				svc, _, bc := newCRUDService()
				_, err := svc.ReviewSubmission(context.Background(), tt.mentorID, 1, tt.materialID, tt.targetUserID, tt.req)
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("want %v, got %v", tt.wantErr, err)
				}
				if len(bc.sent) != 0 {
					t.Errorf("tidak boleh broadcast saat review gagal, got %v", bc.sent)
				}
			})
		}
	})
}