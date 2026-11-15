package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake: pakai ulang fakeRepo + data anggota ─────────────────

type fakeMonitoringRepo struct {
	*fakeRepo                     // GetActiveRole, ListMaterials, GetUserNames diambil dari sini
	members   []entity.MemberInfo // semua anggota cohort 1 (aktif & sudah keluar)
}

func (f *fakeMonitoringRepo) ListActiveMembers(_ context.Context, cohortID uint, search string) ([]entity.MemberInfo, error) {
	var out []entity.MemberInfo
	if cohortID != 1 {
		return out, nil
	}
	for _, m := range f.members {
		if m.Status != "active" || m.Role != entity.RoleMember {
			continue
		}
		if search != "" && !strings.Contains(strings.ToLower(m.Name), strings.ToLower(search)) {
			continue
		}
		out = append(out, m)
	}
	return out, nil
}

func (f *fakeMonitoringRepo) GetMember(_ context.Context, cohortID, userID uint) (*entity.MemberInfo, error) {
	if cohortID == 1 {
		for _, m := range f.members {
			if m.UserID == userID {
				mm := m
				return &mm, nil
			}
		}
	}
	return nil, response.ErrNotFound
}

func (f *fakeMonitoringRepo) ListCohortProgress(_ context.Context, cohortID uint) ([]entity.Progress, error) {
	var out []entity.Progress
	for _, p := range f.progresses {
		if p.CohortID == cohortID {
			out = append(out, p)
		}
	}
	return out, nil
}

// ───────────────── fixture ─────────────────
// Dari newTestService (cohort 1):
//   materi 10 (lewat deadline), 11 (deadline besok), 12 (lewat deadline)
//   Budi (2): 10 submitted, 11 in_progress (ada review), 12 completed
//   Sari (3): 10 completed, 11 & 12 belum mulai
// Tambahan di sini:
//   Tono (4): SUDAH KELUAR, punya progres 11 completed → tidak boleh dihitung di dashboard/matriks

func newMonitoringTest() (*monitoringService, *fakeMonitoringRepo) {
	_, base := newTestService()
	base.progresses = append(base.progresses,
		entity.Progress{ID: 5, CohortID: 1, MaterialID: 11, UserID: 4, Status: entity.StatusCompleted},
	)

	repo := &fakeMonitoringRepo{
		fakeRepo: base,
		members: []entity.MemberInfo{
			{UserID: 2, Name: "Budi", Role: entity.RoleMember, Status: "active"},
			{UserID: 3, Name: "Sari", Role: entity.RoleMember, Status: "active"},
			{UserID: 4, Name: "Tono", Role: entity.RoleMember, Status: "left"},
			{UserID: 1, Name: "Mentor Satu", Role: entity.RoleMentor, Status: "active"},
		},
	}
	svc := &monitoringService{repo: repo, now: func() time.Time { return fixedNow }}
	return svc, repo
}

// ───────────────── Dashboard ─────────────────

func TestDashboard(t *testing.T) {
	svc, _ := newMonitoringTest()

	got, err := svc.Dashboard(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	want := dto.DashboardResponse{
		ActiveMemberCount:  2, // Budi & Sari. Tono keluar, mentor tidak dihitung
		MaterialCount:      3,
		PendingReviewCount: 1,         // Budi materi 10
		CompletedCount:     2,         // Budi 12 + Sari 10. Tono 11 tidak dihitung
		CompletionRate:     2.0 / 6.0, // 2 member × 3 materi
		LateCount:          2,         // Budi 10 (submitted) + Sari 12 (belum mulai)
	}
	if *got != want {
		t.Fatalf("want %+v\ngot  %+v", want, *got)
	}
}

func TestDashboard_TanpaMateri(t *testing.T) {
	svc, repo := newMonitoringTest()
	repo.materials = nil

	got, err := svc.Dashboard(context.Background(), 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if got.MaterialCount != 0 || got.CompletionRate != 0 {
		t.Fatalf("tanpa materi semua ringkasan harus 0, got %+v", *got)
	}
}

func TestDashboard_Ditolak(t *testing.T) {
	tests := []struct {
		name   string
		userID uint
	}{
		{name: "member tidak boleh lihat dashboard", userID: 2},
		{name: "bukan anggota", userID: 9},
		{name: "mantan anggota", userID: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newMonitoringTest()
			_, err := svc.Dashboard(context.Background(), tt.userID, 1)
			if !errors.Is(err, response.ErrForbidden) {
				t.Fatalf("want ErrForbidden, got %v", err)
			}
		})
	}
}

// ───────────────── Matrix ─────────────────

func TestMatrix_Lengkap(t *testing.T) {
	svc, _ := newMonitoringTest()

	got, err := svc.Matrix(context.Background(), 1, 1, dto.MatrixQuery{})
	if err != nil {
		t.Fatal(err)
	}

	if len(got.Materials) != 3 {
		t.Fatalf("want 3 kolom materi, got %d", len(got.Materials))
	}
	if len(got.Rows) != 2 {
		t.Fatalf("want 2 baris (member aktif saja), got %d", len(got.Rows))
	}
	for _, r := range got.Rows {
		if r.UserID == 4 {
			t.Fatal("mantan anggota tidak boleh muncul di matriks")
		}
	}

	budi := got.Rows[0]
	wantStatus := []string{entity.StatusSubmitted, entity.StatusInProgress, entity.StatusCompleted}
	for i, s := range wantStatus {
		if budi.Cells[i].Status != s {
			t.Errorf("Budi sel %d: want %s, got %s", i, s, budi.Cells[i].Status)
		}
	}
	if budi.Completed != 1 || budi.CompletionRate != 1.0/3.0 {
		t.Errorf("Budi: want completed 1 rate 1/3, got %d %v", budi.Completed, budi.CompletionRate)
	}

	sari := got.Rows[1]
	if sari.Cells[1].ProgressID != nil || sari.Cells[1].Status != entity.StatusNotStarted {
		t.Errorf("Sari materi 11 belum mulai: want progress_id nil & not_started, got %+v", sari.Cells[1])
	}
	if !sari.Cells[2].IsLate {
		t.Error("Sari materi 12: lewat deadline & belum mulai → harus terlambat")
	}
}

func TestMatrix_Filter(t *testing.T) {
	tests := []struct {
		name        string
		query       dto.MatrixQuery
		wantUserIDs []uint
		wantCols    int
	}{
		{name: "cari nama peserta", query: dto.MatrixQuery{Q: "sar"}, wantUserIDs: []uint{3}, wantCols: 3},
		{name: "cari nama tidak ada", query: dto.MatrixQuery{Q: "zzz"}, wantUserIDs: []uint{}, wantCols: 3},
		{name: "status submitted", query: dto.MatrixQuery{Status: entity.StatusSubmitted}, wantUserIDs: []uint{2}, wantCols: 3},
		{name: "satu kolom materi", query: dto.MatrixQuery{MaterialID: 11}, wantUserIDs: []uint{2, 3}, wantCols: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newMonitoringTest()
			got, err := svc.Matrix(context.Background(), 1, 1, tt.query)
			if err != nil {
				t.Fatal(err)
			}
			if len(got.Materials) != tt.wantCols {
				t.Errorf("want %d kolom, got %d", tt.wantCols, len(got.Materials))
			}
			if len(got.Rows) != len(tt.wantUserIDs) {
				t.Fatalf("want %d baris, got %d", len(tt.wantUserIDs), len(got.Rows))
			}
			for i, id := range tt.wantUserIDs {
				if got.Rows[i].UserID != id {
					t.Errorf("baris %d: want user %d, got %d", i, id, got.Rows[i].UserID)
				}
			}
		})
	}
}

func TestMatrix_FilterKolomTidakMengubahProgresPribadi(t *testing.T) {
	svc, _ := newMonitoringTest()

	got, err := svc.Matrix(context.Background(), 1, 1, dto.MatrixQuery{MaterialID: 11})
	if err != nil {
		t.Fatal(err)
	}
	// Budi tetap 1 dari 3 materi selesai walau hanya kolom 11 yang tampil
	if got.Rows[0].CompletionRate != 1.0/3.0 {
		t.Errorf("progres pribadi harus dari semua materi, got %v", got.Rows[0].CompletionRate)
	}
}

func TestMatrix_Ditolak(t *testing.T) {
	tests := []struct {
		name    string
		userID  uint
		query   dto.MatrixQuery
		wantErr error
	}{
		{name: "member tidak boleh lihat matriks", userID: 2, wantErr: response.ErrForbidden},
		{name: "bukan anggota", userID: 9, wantErr: response.ErrForbidden},
		{name: "material_id milik cohort lain", userID: 1, query: dto.MatrixQuery{MaterialID: 20}, wantErr: response.ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newMonitoringTest()
			_, err := svc.Matrix(context.Background(), tt.userID, 1, tt.query)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
		})
	}
}

// ───────────────── MemberProgress ─────────────────

func TestMemberProgress_MentorLihatMember(t *testing.T) {
	svc, _ := newMonitoringTest()

	got, err := svc.MemberProgress(context.Background(), 1, 1, 2)
	if err != nil {
		t.Fatal(err)
	}

	if got.Member.Name != "Budi" || got.Member.Status != "active" {
		t.Errorf("member salah: %+v", got.Member)
	}
	wantSummary := dto.PersonalSummary{
		TotalMaterials: 3, NotStarted: 0, InProgress: 1, Submitted: 1, Completed: 1, Late: 1, CompletionRate: 1.0 / 3.0,
	}
	if got.Summary != wantSummary {
		t.Errorf("summary: want %+v\ngot  %+v", wantSummary, got.Summary)
	}
	if len(got.Items) != 3 {
		t.Fatalf("want 3 item, got %d", len(got.Items))
	}

	// materi 11: Mentor bisa lihat catatan review
	review := got.Items[1].Progress.LastReview
	if review == nil || review.Note == nil || review.Reviewer.Name != "Mentor Satu" {
		t.Errorf("mentor harus bisa lihat catatan review lengkap, got %+v", review)
	}
	// materi 10: Mentor bisa lihat isi pengumpulan
	if got.Items[0].Progress.Submission == nil {
		t.Error("mentor harus bisa lihat isi pengumpulan")
	}
}

func TestMemberProgress_MantanAnggotaTetapBisaDilihatMentor(t *testing.T) {
	svc, _ := newMonitoringTest()

	got, err := svc.MemberProgress(context.Background(), 1, 1, 4)
	if err != nil {
		t.Fatal(err)
	}
	if got.Member.Status != "left" {
		t.Errorf("want status left, got %s", got.Member.Status)
	}
	if got.Summary.Completed != 1 {
		t.Errorf("progres mantan anggota harus tetap tercatat, got completed %d", got.Summary.Completed)
	}
}

func TestMemberProgress_MemberLihatMiliknyaSendiri(t *testing.T) {
	svc, _ := newMonitoringTest()

	got, err := svc.MemberProgress(context.Background(), 2, 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if got.Member.UserID != 2 {
		t.Errorf("want user 2, got %d", got.Member.UserID)
	}
}

func TestMemberProgress_Ditolak(t *testing.T) {
	tests := []struct {
		name     string
		userID   uint
		memberID uint
		wantErr  error
	}{
		{name: "member lihat hasil member lain", userID: 2, memberID: 3, wantErr: response.ErrForbidden},
		{name: "bukan anggota", userID: 9, memberID: 2, wantErr: response.ErrForbidden},
		{name: "user tidak pernah jadi anggota cohort", userID: 1, memberID: 99, wantErr: response.ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newMonitoringTest()
			_, err := svc.MemberProgress(context.Background(), tt.userID, 1, tt.memberID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
		})
	}
}
