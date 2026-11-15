package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake repository (data di memori, tanpa DB) ─────────────────

type fakeRepo struct {
	roles      map[[2]uint]string // key: [cohortID, userID]
	materials  []entity.Material
	progresses []entity.Progress
	lastUpdate map[string]any // dicatat supaya test bisa cek field apa yang dikirim ke DB

}

func (f *fakeRepo) GetActiveRole(_ context.Context, cohortID, userID uint) (string, error) {
	if r, ok := f.roles[[2]uint{cohortID, userID}]; ok {
		return r, nil
	}
	return "", response.ErrForbidden
}

func (f *fakeRepo) ListMaterials(_ context.Context, cohortID uint, _ string) ([]entity.Material, error) {
	var out []entity.Material
	for _, m := range f.materials {
		if m.CohortID == cohortID {
			out = append(out, m)
		}
	}
	return out, nil
}

func (f *fakeRepo) GetMaterial(_ context.Context, cohortID, materialID uint) (*entity.Material, error) {
	for _, m := range f.materials {
		if m.ID == materialID && m.CohortID == cohortID {
			mm := m
			return &mm, nil
		}
	}
	return nil, response.ErrNotFound
}

func (f *fakeRepo) ListUserProgress(_ context.Context, cohortID, userID uint) ([]entity.Progress, error) {
	var out []entity.Progress
	for _, p := range f.progresses {
		if p.CohortID == cohortID && p.UserID == userID {
			out = append(out, p)
		}
	}
	return out, nil
}

func (f *fakeRepo) GetUserProgress(_ context.Context, materialID, userID uint) (*entity.Progress, error) {
	for _, p := range f.progresses {
		if p.MaterialID == materialID && p.UserID == userID {
			pp := p
			return &pp, nil
		}
	}
	return nil, nil
}

func (f *fakeRepo) GetUserNames(_ context.Context, ids []uint) (map[uint]string, error) {
	out := map[uint]string{}
	for _, id := range ids {
		if id == 1 {
			out[id] = "Mentor Satu"
		}
	}
	return out, nil
}

func (f *fakeRepo) NextPosition(_ context.Context, cohortID uint) (int, error) {
	maxPos := 0
	for _, m := range f.materials {
		if m.CohortID == cohortID && m.Position > maxPos {
			maxPos = m.Position
		}
	}
	return maxPos + 1, nil
}

func (f *fakeRepo) CreateMaterial(_ context.Context, m *entity.Material) error {
	m.ID = uint(100 + len(f.materials))
	f.materials = append(f.materials, *m)
	return nil
}

func (f *fakeRepo) UpdateMaterial(_ context.Context, cohortID, materialID uint, fields map[string]any) error {
	for _, m := range f.materials {
		if m.ID == materialID && m.CohortID == cohortID {
			f.lastUpdate = fields
			return nil
		}
	}
	return response.ErrNotFound
}

func (f *fakeRepo) DeleteMaterial(_ context.Context, cohortID, materialID uint) error {
	for i, m := range f.materials {
		if m.ID == materialID && m.CohortID == cohortID {
			f.materials = append(f.materials[:i], f.materials[i+1:]...)
			return nil
		}
	}
	return response.ErrNotFound
}

func (f *fakeRepo) ReorderMaterials(context.Context, uint, []uint) error {
	return nil
}

// ───────────────── fixture ─────────────────
// cohort 1: user 1 = mentor, user 2 & 3 = member. user 9 bukan anggota.
// materi 10 (lewat deadline), 11 (deadline besok), 12 (lewat deadline), 20 milik cohort 2.

var fixedNow = time.Date(2026, 11, 15, 12, 0, 0, 0, time.UTC)

func newTestService() (*materialService, *fakeRepo) {
	past := fixedNow.Add(-24 * time.Hour)
	future := fixedNow.Add(24 * time.Hour)
	reviewedAt := fixedNow.Add(-time.Hour)
	submittedAt := fixedNow.Add(-2 * time.Hour)
	reviewer := uint(1)
	decision := "revision_requested"
	note := "tambahkan error handling"
	content := "sudah bikin endpoint"

	repo := &fakeRepo{
		roles: map[[2]uint]string{
			{1, 1}: entity.RoleMentor,
			{1, 2}: entity.RoleMember,
			{1, 3}: entity.RoleMember,
		},
		materials: []entity.Material{
			{ID: 10, CohortID: 1, Title: "Gin", Position: 1, Deadline: &past},
			{ID: 11, CohortID: 1, Title: "Gorm", Position: 2, Deadline: &future},
			{ID: 12, CohortID: 1, Title: "Redis", Position: 3, Deadline: &past},
			{ID: 20, CohortID: 2, Title: "Cohort lain", Position: 1},
		},
		progresses: []entity.Progress{
			{ID: 1, CohortID: 1, MaterialID: 10, UserID: 2, Status: entity.StatusSubmitted,
				SubmissionContent: &content, SubmittedAt: &submittedAt},
			{ID: 2, CohortID: 1, MaterialID: 12, UserID: 2, Status: entity.StatusCompleted},
			{ID: 3, CohortID: 1, MaterialID: 11, UserID: 2, Status: entity.StatusInProgress,
				ReviewDecision: &decision, ReviewNote: &note, ReviewerID: &reviewer, ReviewedAt: &reviewedAt},
			{ID: 4, CohortID: 1, MaterialID: 10, UserID: 3, Status: entity.StatusCompleted}, // milik member lain
		},
	}
	svc := &materialService{repo: repo, rt: contract.NoopBroadcaster{}, now: func() time.Time { return fixedNow }}
	return svc, repo
}

func boolPtr(b bool) *bool { return &b }

// ───────────────── ListMaterials ─────────────────

func TestListMaterials(t *testing.T) {
	tests := []struct {
		name         string
		userID       uint
		query        dto.ListMaterialsQuery
		wantErr      error
		wantIDs      []uint
		wantNoMyProg bool
	}{
		{name: "member: semua materi urut sesuai position", userID: 2, wantIDs: []uint{10, 11, 12}},
		{name: "filter status submitted", userID: 2, query: dto.ListMaterialsQuery{Status: entity.StatusSubmitted}, wantIDs: []uint{10}},
		{name: "filter in_progress", userID: 2, query: dto.ListMaterialsQuery{Status: entity.StatusInProgress}, wantIDs: []uint{11}},
		{name: "member lain: tanpa row progres = not_started", userID: 3, query: dto.ListMaterialsQuery{Status: entity.StatusNotStarted}, wantIDs: []uint{11, 12}},
		{name: "late=true: lewat deadline & belum completed", userID: 2, query: dto.ListMaterialsQuery{Late: boolPtr(true)}, wantIDs: []uint{10}},
		{name: "late=false", userID: 2, query: dto.ListMaterialsQuery{Late: boolPtr(false)}, wantIDs: []uint{11, 12}},
		{name: "mentor: tanpa my_progress, filter diabaikan", userID: 1, query: dto.ListMaterialsQuery{Status: entity.StatusSubmitted}, wantIDs: []uint{10, 11, 12}, wantNoMyProg: true},
		{name: "bukan anggota: ditolak", userID: 9, wantErr: response.ErrForbidden},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newTestService()
			got, err := svc.ListMaterials(context.Background(), tt.userID, 1, tt.query)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("want err %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if len(got) != len(tt.wantIDs) {
				t.Fatalf("want %d materi, got %d", len(tt.wantIDs), len(got))
			}
			for i, id := range tt.wantIDs {
				if got[i].ID != id {
					t.Errorf("index %d: want id %d, got %d", i, id, got[i].ID)
				}
				if tt.wantNoMyProg && got[i].MyProgress != nil {
					t.Errorf("mentor seharusnya my_progress = nil")
				}
				if !tt.wantNoMyProg && got[i].MyProgress == nil {
					t.Errorf("member seharusnya punya my_progress")
				}
			}
		})
	}
}

func TestListMaterials_ProgresMemberLainTidakBocor(t *testing.T) {
	svc, _ := newTestService()

	got, err := svc.ListMaterials(context.Background(), 2, 1, dto.ListMaterialsQuery{})
	if err != nil {
		t.Fatal(err)
	}
	// materi 10: member 3 completed, member 2 harus tetap lihat miliknya (submitted)
	if got[0].MyProgress.Status != entity.StatusSubmitted {
		t.Fatalf("want submitted (milik sendiri), got %s", got[0].MyProgress.Status)
	}
}

// ───────────────── GetMaterial ─────────────────

func TestGetMaterial_MemberLihatProgresDanReviewSendiri(t *testing.T) {
	svc, _ := newTestService()

	got, err := svc.GetMaterial(context.Background(), 2, 1, 11)
	if err != nil {
		t.Fatal(err)
	}
	p := got.MyProgress
	if p == nil || p.Status != entity.StatusInProgress {
		t.Fatalf("want in_progress, got %+v", p)
	}
	if p.LastReview == nil || p.LastReview.Decision != "revision_requested" {
		t.Fatalf("want last_review revision_requested, got %+v", p.LastReview)
	}
	if p.LastReview.Reviewer.Name != "Mentor Satu" {
		t.Errorf("want nama reviewer terisi, got %q", p.LastReview.Reviewer.Name)
	}
}

func TestGetMaterial_SubmissionTerlambat(t *testing.T) {
	svc, _ := newTestService()

	got, err := svc.GetMaterial(context.Background(), 2, 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	if got.MyProgress.Submission == nil {
		t.Fatal("want submission terisi")
	}
	// deadline kemarin, submit 2 jam lalu → terlambat, tapi tetap diterima
	if !got.MyProgress.Submission.IsLate {
		t.Error("want submission.is_late = true")
	}
}

func TestGetMaterial_Ditolak(t *testing.T) {
	tests := []struct {
		name       string
		userID     uint
		materialID uint
		wantErr    error
	}{
		{name: "materi milik cohort lain", userID: 2, materialID: 20, wantErr: response.ErrNotFound},
		{name: "materi tidak ada", userID: 2, materialID: 999, wantErr: response.ErrNotFound},
		{name: "bukan anggota cohort", userID: 9, materialID: 10, wantErr: response.ErrForbidden},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newTestService()
			_, err := svc.GetMaterial(context.Background(), tt.userID, 1, tt.materialID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestGetMaterial_MentorTanpaMyProgress(t *testing.T) {
	svc, _ := newTestService()

	got, err := svc.GetMaterial(context.Background(), 1, 1, 10)
	if err != nil {
		t.Fatal(err)
	}
	if got.MyProgress != nil {
		t.Errorf("mentor seharusnya my_progress = nil")
	}
}

// ───────────────── MySummary ─────────────────

func TestMySummary(t *testing.T) {
	svc, _ := newTestService()

	got, err := svc.MySummary(context.Background(), 2, 1)
	if err != nil {
		t.Fatal(err)
	}
	want := dto.PersonalSummary{
		TotalMaterials: 3,
		NotStarted:     0,
		InProgress:     1, // materi 11
		Submitted:      1, // materi 10
		Completed:      1, // materi 12
		Late:           1, // materi 10 (lewat deadline & belum completed). Materi 12 completed → tidak dihitung
		CompletionRate: 1.0 / 3.0,
	}
	if *got != want {
		t.Fatalf("want %+v\ngot  %+v", want, *got)
	}
}

func TestMySummary_TanpaMateri(t *testing.T) {
	svc, repo := newTestService()
	repo.materials = nil

	got, err := svc.MySummary(context.Background(), 2, 1)
	if err != nil {
		t.Fatal(err)
	}
	if got.TotalMaterials != 0 || got.CompletionRate != 0 {
		t.Fatalf("want semua 0, got %+v", *got)
	}
}

func TestMySummary_Ditolak(t *testing.T) {
	tests := []struct {
		name   string
		userID uint
	}{
		{name: "mentor tidak punya progres pribadi", userID: 1},
		{name: "bukan anggota", userID: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _ := newTestService()
			_, err := svc.MySummary(context.Background(), tt.userID, 1)
			if !errors.Is(err, response.ErrForbidden) {
				t.Fatalf("want ErrForbidden, got %v", err)
			}
		})
	}
}

// ───────────────── fake broadcaster: catat event yang dikirim ─────────────────

type sentEvent struct{ room, event string }

type fakeBroadcaster struct{ sent []sentEvent }

func (b *fakeBroadcaster) Broadcast(room, event string, _ any) {
	b.sent = append(b.sent, sentEvent{room, event})
}

func (b *fakeBroadcaster) has(room, event string) bool {
	for _, e := range b.sent {
		if e.room == room && e.event == event {
			return true
		}
	}
	return false
}

func newCRUDService() (*materialService, *fakeRepo, *fakeBroadcaster) {
	svc, repo := newTestService()
	bc := &fakeBroadcaster{}
	svc.rt = bc
	return svc, repo, bc
}

func strPtr(s string) *string { return &s }

// ───────────────── CreateMaterial ─────────────────

func TestCreateMaterial_Mentor(t *testing.T) {
	svc, _, bc := newCRUDService()
	wib := time.FixedZone("WIB", 7*3600)
	deadline := time.Date(2026, 11, 20, 23, 59, 0, 0, wib)

	got, err := svc.CreateMaterial(context.Background(), 1, 1, dto.CreateMaterialRequest{
		Title:    "  Pengenalan Gin  ",
		Guide:    "baca docs routing",
		Deadline: &deadline,
	})
	if err != nil {
		t.Fatal(err)
	}

	if got.Position != 4 {
		t.Errorf("materi baru harus di urutan terakhir (4), got %d", got.Position)
	}
	if got.Title != "Pengenalan Gin" {
		t.Errorf("title harus di-trim, got %q", got.Title)
	}
	if got.References == nil {
		t.Error("references harus [] bukan nil")
	}
	if got.Deadline == nil || got.Deadline.Location() != time.UTC || !got.Deadline.Equal(deadline) {
		t.Errorf("deadline harus disimpan UTC dengan waktu yang sama, got %v", got.Deadline)
	}
	if got.MyProgress != nil {
		t.Error("response untuk mentor tidak boleh berisi my_progress")
	}
	if !bc.has("cohort:1", "material.created") {
		t.Error("harus broadcast material.created ke room cohort")
	}
	if !bc.has("cohort:1:mentors", "dashboard.updated") {
		t.Error("harus broadcast dashboard.updated ke room mentor")
	}
}

func TestCreateMaterial_Ditolak(t *testing.T) {
	tests := []struct {
		name    string
		userID  uint
		req     dto.CreateMaterialRequest
		wantErr error
	}{
		{name: "member tidak boleh buat materi", userID: 2, req: dto.CreateMaterialRequest{Title: "Gin", Guide: "x"}, wantErr: response.ErrForbidden},
		{name: "bukan anggota", userID: 9, req: dto.CreateMaterialRequest{Title: "Gin", Guide: "x"}, wantErr: response.ErrForbidden},
		{name: "title cuma spasi + 2 huruf", userID: 1, req: dto.CreateMaterialRequest{Title: "   ab ", Guide: "x"}, wantErr: response.ErrBadRequest},
		{name: "guide cuma spasi", userID: 1, req: dto.CreateMaterialRequest{Title: "Gin", Guide: "   "}, wantErr: response.ErrBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, repo, bc := newCRUDService()
			before := len(repo.materials)

			_, err := svc.CreateMaterial(context.Background(), tt.userID, 1, tt.req)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
			if len(repo.materials) != before {
				t.Error("materi tidak boleh tersimpan kalau ditolak")
			}
			if len(bc.sent) != 0 {
				t.Errorf("tidak boleh ada broadcast kalau gagal, got %v", bc.sent)
			}
		})
	}
}

// ───────────────── UpdateMaterial ─────────────────

func TestUpdateMaterial_HapusDeadline(t *testing.T) {
	svc, repo, bc := newCRUDService()

	_, err := svc.UpdateMaterial(context.Background(), 1, 1, 10, dto.UpdateMaterialRequest{
		Deadline: dto.NullableTime{Set: true, Value: nil},
	})
	if err != nil {
		t.Fatal(err)
	}

	v, ok := repo.lastUpdate["deadline"]
	if !ok || v != nil {
		t.Fatalf("deadline:null harus jadi field deadline=nil, got ok=%v v=%v", ok, v)
	}
	if !bc.has("cohort:1", "material.updated") {
		t.Error("harus broadcast material.updated")
	}
}

func TestUpdateMaterial_FieldYangTidakDikirimTidakDiubah(t *testing.T) {
	svc, repo, _ := newCRUDService()

	_, err := svc.UpdateMaterial(context.Background(), 1, 1, 10, dto.UpdateMaterialRequest{
		Title: strPtr("  Gin Lanjutan "),
	})
	if err != nil {
		t.Fatal(err)
	}

	if repo.lastUpdate["title"] != "Gin Lanjutan" {
		t.Errorf("title harus di-trim, got %v", repo.lastUpdate["title"])
	}
	for _, key := range []string{"deadline", "guide", "references"} {
		if _, ok := repo.lastUpdate[key]; ok {
			t.Errorf("field %q tidak dikirim, jadi tidak boleh ikut di-update", key)
		}
	}
}

func TestUpdateMaterial_Ditolak(t *testing.T) {
	tests := []struct {
		name       string
		userID     uint
		materialID uint
		req        dto.UpdateMaterialRequest
		wantErr    error
	}{
		{name: "tidak ada field", userID: 1, materialID: 10, req: dto.UpdateMaterialRequest{}, wantErr: response.ErrBadRequest},
		{name: "title terlalu pendek", userID: 1, materialID: 10, req: dto.UpdateMaterialRequest{Title: strPtr(" a ")}, wantErr: response.ErrBadRequest},
		{name: "member tidak boleh edit", userID: 2, materialID: 10, req: dto.UpdateMaterialRequest{Title: strPtr("Gin 2")}, wantErr: response.ErrForbidden},
		{name: "materi milik cohort lain", userID: 1, materialID: 20, req: dto.UpdateMaterialRequest{Title: strPtr("Gin 2")}, wantErr: response.ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, _, bc := newCRUDService()
			_, err := svc.UpdateMaterial(context.Background(), tt.userID, 1, tt.materialID, tt.req)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
			if len(bc.sent) != 0 {
				t.Errorf("tidak boleh ada broadcast kalau gagal, got %v", bc.sent)
			}
		})
	}
}

// ───────────────── DeleteMaterial ─────────────────

func TestDeleteMaterial_Mentor(t *testing.T) {
	svc, repo, bc := newCRUDService()

	if err := svc.DeleteMaterial(context.Background(), 1, 1, 10); err != nil {
		t.Fatal(err)
	}
	for _, m := range repo.materials {
		if m.ID == 10 {
			t.Fatal("materi 10 harus terhapus")
		}
	}
	if !bc.has("cohort:1", "material.deleted") {
		t.Error("harus broadcast material.deleted")
	}
	if !bc.has("cohort:1:mentors", "dashboard.updated") {
		t.Error("hapus materi mengubah dashboard, harus broadcast dashboard.updated")
	}
}

func TestDeleteMaterial_Ditolak(t *testing.T) {
	tests := []struct {
		name       string
		userID     uint
		materialID uint
		wantErr    error
	}{
		{name: "member tidak boleh hapus", userID: 2, materialID: 10, wantErr: response.ErrForbidden},
		{name: "materi milik cohort lain", userID: 1, materialID: 20, wantErr: response.ErrNotFound},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc, repo, bc := newCRUDService()
			before := len(repo.materials)

			err := svc.DeleteMaterial(context.Background(), tt.userID, 1, tt.materialID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v, got %v", tt.wantErr, err)
			}
			if len(repo.materials) != before {
				t.Error("tidak boleh ada materi terhapus kalau ditolak")
			}
			if len(bc.sent) != 0 {
				t.Errorf("tidak boleh ada broadcast, got %v", bc.sent)
			}
		})
	}
}

// ───────────────── ReorderMaterials ─────────────────

func TestReorderMaterials_Mentor(t *testing.T) {
	svc, _, bc := newCRUDService()

	got, err := svc.ReorderMaterials(context.Background(), 1, 1, dto.ReorderMaterialsRequest{MaterialIDs: []uint{12, 10, 11}})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 3 {
		t.Errorf("want 3 materi dikembalikan, got %d", len(got))
	}
	if !bc.has("cohort:1", "material.reordered") {
		t.Error("harus broadcast material.reordered")
	}
}

func TestReorderMaterials_MemberDitolak(t *testing.T) {
	svc, _, bc := newCRUDService()

	_, err := svc.ReorderMaterials(context.Background(), 2, 1, dto.ReorderMaterialsRequest{MaterialIDs: []uint{12, 10, 11}})
	if !errors.Is(err, response.ErrForbidden) {
		t.Fatalf("want ErrForbidden, got %v", err)
	}
	if len(bc.sent) != 0 {
		t.Errorf("tidak boleh ada broadcast, got %v", bc.sent)
	}
}
