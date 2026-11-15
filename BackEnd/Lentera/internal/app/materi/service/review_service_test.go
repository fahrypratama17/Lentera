package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake: pakai ulang fakeMonitoringRepo ─────────────────

type fakeReviewRepo struct {
	*fakeMonitoringRepo
	// kalau diisi: simulasi mentor lain sudah mengubah status tepat sebelum SaveReview
	concurrentStatus string
}

func (f *fakeReviewRepo) ListCohortMembers(_ context.Context, cohortID uint) ([]entity.MemberInfo, error) {
	if cohortID != 1 {
		return nil, nil
	}
	return f.members, nil
}

func (f *fakeReviewRepo) ListProgressByStatus(_ context.Context, cohortID uint, status string, materialID uint) ([]entity.Progress, error) {
	var out []entity.Progress
	for _, p := range f.progresses {
		if p.CohortID != cohortID || p.Status != status {
			continue
		}
		if materialID != 0 && p.MaterialID != materialID {
			continue
		}
		out = append(out, p)
	}
	return out, nil
}

func (f *fakeReviewRepo) GetProgress(_ context.Context, cohortID, progressID uint) (*entity.Progress, error) {
	for _, p := range f.progresses {
		if p.ID == progressID && p.CohortID == cohortID {
			pp := p // salinan: perubahan di service belum menyentuh "DB"
			return &pp, nil
		}
	}
	return nil, response.ErrNotFound
}

// meniru UPDATE ... WHERE status = fromStatus
func (f *fakeReviewRepo) SaveReview(_ context.Context, p *entity.Progress, fromStatus string) error {
	for i := range f.progresses {
		if f.progresses[i].ID == p.ID && f.progresses[i].CohortID == p.CohortID {
			if f.concurrentStatus != "" {
				f.progresses[i].Status = f.concurrentStatus
			}
			if f.progresses[i].Status != fromStatus {
				return fmt.Errorf("%w: progres sudah direview", response.ErrConflict)
			}
			f.progresses[i] = *p
			return nil
		}
	}
	return fmt.Errorf("%w: progres tidak ada", response.ErrConflict)
}

func (f *fakeReviewRepo) find(id uint) entity.Progress {
	for _, p := range f.progresses {
		if p.ID == id {
			return p
		}
	}
	return entity.Progress{}
}

// ───────────────── fixture ─────────────────
// Dari newMonitoringTest (cohort 1): Mentor(1), Budi(2), Sari(3), Tono(4, SUDAH KELUAR)
//   progres 1: Budi  m10 submitted
//   progres 2: Budi  m12 completed
//   progres 3: Budi  m11 in_progress (ada review sebelumnya)
//   progres 4: Sari  m10 completed
//   progres 5: Tono  m11 completed
// Tambahan:
//   progres 6: Tono  m10 submitted   ← mantan anggota
//   progres 7: Sari  m11 submitted

func newReviewTest() (*reviewService, *fakeReviewRepo, *fakeBroadcaster) {
	_, mon := newMonitoringTest()
	sub := fixedNow.Add(-30 * time.Minute)
	mon.progresses = append(mon.progresses,
		entity.Progress{ID: 6, CohortID: 1, MaterialID: 10, UserID: 4, Status: entity.StatusSubmitted},
		entity.Progress{ID: 7, CohortID: 1, MaterialID: 11, UserID: 3, Status: entity.StatusSubmitted,
			SubmissionContent: strPtr("hasil sari"), SubmittedAt: &sub},
	)

	repo := &fakeReviewRepo{fakeMonitoringRepo: mon}
	bc := &fakeBroadcaster{}
	svc := &reviewService{repo: repo, rt: bc, now: func() time.Time { return fixedNow }}
	return svc, repo, bc
}

func progressIDs(items []dto.ReviewItemResponse) []uint {
	out := make([]uint, 0, len(items))
	for _, it := range items {
		out = append(out, *it.ID)
	}
	return out
}

func equalIDs(a, b []uint) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// ───────────────── ListReviews ─────────────────

func TestListReviews(t *testing.T) {
	tests := []struct {
		name    string
		query   dto.ListReviewsQuery
		wantIDs []uint
	}{
		{name: "default: submitted, member aktif saja", query: dto.ListReviewsQuery{}, wantIDs: []uint{1, 7}},
		{name: "include_inactive: mantan anggota ikut", query: dto.ListReviewsQuery{IncludeInactive: true}, wantIDs: []uint{1, 6, 7}},
		{name: "filter materi", query: dto.ListReviewsQuery{MaterialID: 11}, wantIDs: []uint{7}},
		{name: "cari nama member", query: dto.ListReviewsQuery{Q: "bud"}, wantIDs: []uint{1}},
		{name: "status completed", query: dto.ListReviewsQuery{Status: entity.StatusCompleted}, wantIDs: []uint{2, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _, _ := newReviewTest()
			got, err := svc.ListReviews(context.Background(), 1, 1, tt.query)
			if err != nil {
				t.Fatal(err)
			}
			if ids := progressIDs(got); !equalIDs(ids, tt.wantIDs) {
				t.Fatalf("want %v, got %v", tt.wantIDs, ids)
			}
		})
	}
}

func TestListReviews_IsiItem(t *testing.T) {
	svc, _, _ := newReviewTest()
	got, err := svc.ListReviews(context.Background(), 1, 1, dto.ListReviewsQuery{})
	if err != nil {
		t.Fatal(err)
	}
	budi := got[0]
	if budi.Member.Name != "Budi" || budi.Material.Title != "Gin" {
		t.Errorf("member/material salah: %+v %+v", budi.Member, budi.Material)
	}
	if budi.Submission == nil {
		t.Error("mentor harus bisa lihat isi pengumpulan")
	}
}

func TestListReviews_MemberDitolak(t *testing.T) {
	svc, _, _ := newReviewTest()
	if _, err := svc.ListReviews(context.Background(), 2, 1, dto.ListReviewsQuery{}); !errors.Is(err, response.ErrForbidden) {
		t.Fatalf("want ErrForbidden, got %v", err)
	}
}

// ───────────────── GetReviewItem ─────────────────

func TestGetReviewItem(t *testing.T) {
	tests := []struct {
		name       string
		userID     uint
		progressID uint
		wantErr    error
	}{
		{name: "mentor lihat progres member", userID: 1, progressID: 3},
		{name: "member lihat miliknya sendiri", userID: 2, progressID: 1},
		{name: "member lihat milik member lain", userID: 2, progressID: 7, wantErr: response.ErrForbidden},
		{name: "progres tidak ada", userID: 1, progressID: 999, wantErr: response.ErrNotFound},
		{name: "bukan anggota", userID: 9, progressID: 1, wantErr: response.ErrForbidden},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _, _ := newReviewTest()
			got, err := svc.GetReviewItem(context.Background(), tt.userID, 1, tt.progressID)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("want %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if *got.ID != tt.progressID {
				t.Errorf("want progres %d, got %d", tt.progressID, *got.ID)
			}
		})
	}
}

func TestGetReviewItem_CatatanReviewTerlihatMentor(t *testing.T) {
	svc, _, _ := newReviewTest()
	got, err := svc.GetReviewItem(context.Background(), 1, 1, 3)
	if err != nil {
		t.Fatal(err)
	}
	if got.LastReview == nil || got.LastReview.Reviewer.Name != "Mentor Satu" {
		t.Errorf("catatan review & nama reviewer harus terlihat, got %+v", got.LastReview)
	}
}

// ───────────────── Approve ─────────────────

func TestApprove(t *testing.T) {
	svc, repo, bc := newReviewTest()

	got, err := svc.Approve(context.Background(), 1, 1, 1, dto.ApproveRequest{Note: strPtr("  mantap  ")})
	if err != nil {
		t.Fatal(err)
	}

	saved := repo.find(1)
	if saved.Status != entity.StatusCompleted || saved.CompletedAt == nil {
		t.Errorf("harus completed dengan completed_at, got %s %v", saved.Status, saved.CompletedAt)
	}
	if saved.ReviewNote == nil || *saved.ReviewNote != "mantap" {
		t.Errorf("catatan harus di-trim, got %v", saved.ReviewNote)
	}
	if got.Status != entity.StatusCompleted {
		t.Errorf("response harus status terbaru, got %s", got.Status)
	}

	// keputusan hanya berlaku untuk progres ini
	if repo.find(6).Status != entity.StatusSubmitted {
		t.Error("progres member lain di materi yang sama tidak boleh berubah")
	}

	if !bc.has("user:2", "review.decided") {
		t.Error("member pemilik harus menerima review.decided")
	}
	if !bc.has("cohort:1:mentors", "review.decided") || !bc.has("cohort:1:mentors", "dashboard.updated") {
		t.Error("mentor harus menerima review.decided & dashboard.updated")
	}
	if bc.has("user:3", "review.decided") {
		t.Error("member lain tidak boleh menerima keputusan milik orang lain")
	}
}

func TestApprove_CatatanKosongJadiNil(t *testing.T) {
	svc, repo, _ := newReviewTest()
	if _, err := svc.Approve(context.Background(), 1, 1, 1, dto.ApproveRequest{Note: strPtr("   ")}); err != nil {
		t.Fatal(err)
	}
	if repo.find(1).ReviewNote != nil {
		t.Error("catatan berisi spasi saja harus disimpan sebagai null")
	}
}

// ───────────────── RequestRevision ─────────────────

func TestRequestRevision(t *testing.T) {
	svc, repo, bc := newReviewTest()

	_, err := svc.RequestRevision(context.Background(), 1, 1, 7, dto.ReviewNoteRequest{Note: "tambah error handling"})
	if err != nil {
		t.Fatal(err)
	}

	saved := repo.find(7)
	if saved.Status != entity.StatusInProgress {
		t.Errorf("revisi harus kembali in_progress, got %s", saved.Status)
	}
	if saved.ReviewDecision == nil || *saved.ReviewDecision != entity.DecisionRevisionRequested {
		t.Errorf("decision salah: %v", saved.ReviewDecision)
	}
	if saved.SubmissionContent == nil || *saved.SubmissionContent != "hasil sari" {
		t.Error("isi pengumpulan sebelumnya harus tetap ada setelah diminta revisi")
	}
	if !bc.has("user:3", "review.decided") {
		t.Error("member pemilik harus menerima review.decided")
	}
}

// ───────────────── Reopen ─────────────────

func TestReopen(t *testing.T) {
	svc, repo, _ := newReviewTest()

	_, err := svc.Reopen(context.Background(), 1, 1, 2, dto.ReviewNoteRequest{Note: "ada requirement baru"})
	if err != nil {
		t.Fatal(err)
	}
	saved := repo.find(2)
	if saved.Status != entity.StatusInProgress || saved.CompletedAt != nil {
		t.Errorf("reopen harus in_progress & completed_at kosong, got %s %v", saved.Status, saved.CompletedAt)
	}
}

// ───────────────── keputusan yang ditolak ─────────────────

func TestKeputusanReview_Ditolak(t *testing.T) {
	tests := []struct {
		name       string
		progressID uint
		userID     uint
		act        func(s *reviewService, userID, progressID uint) error
		race       string
		wantErr    error
		wantStatus string // status progres di "DB" harus tetap ini
	}{
		{
			name: "member tidak boleh approve", progressID: 1, userID: 2,
			act: func(s *reviewService, u, p uint) error {
				_, e := s.Approve(context.Background(), u, 1, p, dto.ApproveRequest{})
				return e
			},
			wantErr: response.ErrForbidden, wantStatus: entity.StatusSubmitted,
		},
		{
			name: "approve progres in_progress", progressID: 3, userID: 1,
			act: func(s *reviewService, u, p uint) error {
				_, e := s.Approve(context.Background(), u, 1, p, dto.ApproveRequest{})
				return e
			},
			wantErr: response.ErrConflict, wantStatus: entity.StatusInProgress,
		},
		{
			name: "reopen progres submitted", progressID: 1, userID: 1,
			act: func(s *reviewService, u, p uint) error {
				_, e := s.Reopen(context.Background(), u, 1, p, dto.ReviewNoteRequest{Note: "x"})
				return e
			},
			wantErr: response.ErrConflict, wantStatus: entity.StatusSubmitted,
		},
		{
			name: "revisi tanpa catatan", progressID: 1, userID: 1,
			act: func(s *reviewService, u, p uint) error {
				_, e := s.RequestRevision(context.Background(), u, 1, p, dto.ReviewNoteRequest{Note: "   "})
				return e
			},
			wantErr: response.ErrBadRequest, wantStatus: entity.StatusSubmitted,
		},
		{
			name: "progres tidak ada", progressID: 999, userID: 1,
			act: func(s *reviewService, u, p uint) error {
				_, e := s.Approve(context.Background(), u, 1, p, dto.ApproveRequest{})
				return e
			},
			wantErr: response.ErrNotFound,
		},
		{
			name: "didahului mentor lain (race)", progressID: 1, userID: 1, race: entity.StatusCompleted,
			act: func(s *reviewService, u, p uint) error {
				_, e := s.Approve(context.Background(), u, 1, p, dto.ApproveRequest{})
				return e
			},
			wantErr: response.ErrConflict, wantStatus: entity.StatusCompleted,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, repo, bc := newReviewTest()
			repo.concurrentStatus = tt.race

			err := tt.act(svc, tt.userID, tt.progressID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
			if tt.wantStatus != "" && repo.find(tt.progressID).Status != tt.wantStatus {
				t.Errorf("status di DB harus tetap %s, got %s", tt.wantStatus, repo.find(tt.progressID).Status)
			}
			if len(bc.sent) != 0 {
				t.Errorf("tidak boleh ada broadcast kalau gagal, got %v", bc.sent)
			}
		})
	}
}
