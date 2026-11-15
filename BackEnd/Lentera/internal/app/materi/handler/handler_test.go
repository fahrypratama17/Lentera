package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake service ─────────────────

type fakeService struct {
	err error

	gotUserID   uint
	gotCohortID uint
	gotQuery    dto.ListMaterialsQuery
	gotCreate   dto.CreateMaterialRequest
	gotUpdate   dto.UpdateMaterialRequest
	called      bool // true kalau method CRUD sempat dipanggil
}

func (f *fakeService) ListMaterials(_ context.Context, userID, cohortID uint, q dto.ListMaterialsQuery) ([]dto.MaterialResponse, error) {
	f.gotUserID, f.gotCohortID, f.gotQuery = userID, cohortID, q
	if f.err != nil {
		return nil, f.err
	}
	return []dto.MaterialResponse{{ID: 10, Title: "Gin", References: []string{}}}, nil
}

func (f *fakeService) GetMaterial(_ context.Context, userID, cohortID, materialID uint) (*dto.MaterialResponse, error) {
	f.gotUserID, f.gotCohortID = userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.MaterialResponse{ID: materialID, Title: "Gin", References: []string{}}, nil
}

func (f *fakeService) MySummary(_ context.Context, userID, cohortID uint) (*dto.PersonalSummary, error) {
	f.gotUserID, f.gotCohortID = userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.PersonalSummary{TotalMaterials: 3, Completed: 1, CompletionRate: 1.0 / 3.0}, nil
}

func (f *fakeService) CreateMaterial(_ context.Context, userID, cohortID uint, req dto.CreateMaterialRequest) (*dto.MaterialResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotCreate = true, userID, cohortID, req
	if f.err != nil {
		return nil, f.err
	}
	return &dto.MaterialResponse{ID: 100, CohortID: cohortID, Title: req.Title, References: []string{}}, nil
}

func (f *fakeService) UpdateMaterial(_ context.Context, userID, cohortID, materialID uint, req dto.UpdateMaterialRequest) (*dto.MaterialResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotUpdate = true, userID, cohortID, req
	if f.err != nil {
		return nil, f.err
	}
	return &dto.MaterialResponse{ID: materialID, CohortID: cohortID, Title: "Diubah", References: []string{}}, nil
}

func (f *fakeService) DeleteMaterial(_ context.Context, userID, cohortID, _ uint) error {
	f.called, f.gotUserID, f.gotCohortID = true, userID, cohortID
	return f.err
}

func (f *fakeService) ReorderMaterials(_ context.Context, userID, cohortID uint, _ dto.ReorderMaterialsRequest) ([]dto.MaterialResponse, error) {
	f.called, f.gotUserID, f.gotCohortID = true, userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return []dto.MaterialResponse{}, nil
}

// ───────────────── helper ─────────────────

// uid = 0 artinya simulasi request tanpa login
func newTestRouter(svc *fakeService, uid uint) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// pengganti middleware JWT
	fakeAuth := func(c *gin.Context) {
		if uid != 0 {
			c.Set("uid", uid)
		}
		c.Next()
	}
	New(svc).RegisterRoutes(r.Group("/api/v1"), fakeAuth)
	return r
}

type envelope struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
	Error   *struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func doGET(t *testing.T, r *gin.Engine, path string) (*httptest.ResponseRecorder, envelope) {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var body envelope
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("response bukan JSON valid: %v\nbody: %s", err, w.Body.String())
	}
	return w, body
}

// ───────────────── test: status code & format error ─────────────────

func TestMaterialHandler_StatusCodes(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		uid        uint
		svcErr     error
		wantStatus int
		wantCode   string // kosong = sukses
	}{
		{name: "list sukses", path: "/api/v1/cohorts/1/materials", uid: 2, wantStatus: 200},
		{name: "detail sukses", path: "/api/v1/cohorts/1/materials/10", uid: 2, wantStatus: 200},
		{name: "summary sukses", path: "/api/v1/cohorts/1/me/summary", uid: 2, wantStatus: 200},

		{name: "tanpa login", path: "/api/v1/cohorts/1/materials", uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},
		{name: "cohortId bukan angka", path: "/api/v1/cohorts/abc/materials", uid: 2, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "materialId nol", path: "/api/v1/cohorts/1/materials/0", uid: 2, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "status filter tidak valid", path: "/api/v1/cohorts/1/materials?status=done", uid: 2, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "late bukan boolean", path: "/api/v1/cohorts/1/materials?late=mungkin", uid: 2, wantStatus: 400, wantCode: "BAD_INPUT"},

		{name: "bukan anggota cohort", path: "/api/v1/cohorts/1/materials", uid: 9, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN"},
		{name: "materi tidak ditemukan", path: "/api/v1/cohorts/1/materials/999", uid: 2, svcErr: response.ErrNotFound, wantStatus: 404, wantCode: "NOT_FOUND"},
		{name: "mentor minta summary", path: "/api/v1/cohorts/1/me/summary", uid: 1, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN"},
		{name: "error tak terduga", path: "/api/v1/cohorts/1/materials", uid: 2, svcErr: errors.New("pq: connection refused"), wantStatus: 500, wantCode: "INTERNAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeService{err: tt.svcErr}
			w, body := doGET(t, newTestRouter(svc, tt.uid), tt.path)

			if w.Code != tt.wantStatus {
				t.Fatalf("want status %d, got %d (body: %s)", tt.wantStatus, w.Code, w.Body.String())
			}

			if tt.wantCode == "" {
				if !body.Success || body.Error != nil {
					t.Fatalf("want success, got %s", w.Body.String())
				}
				return
			}

			if body.Success || body.Error == nil {
				t.Fatalf("want error envelope, got %s", w.Body.String())
			}
			if body.Error.Code != tt.wantCode {
				t.Errorf("want code %s, got %s", tt.wantCode, body.Error.Code)
			}
		})
	}
}

// ───────────────── test: detail error internal tidak bocor ─────────────────

func TestMaterialHandler_ErrorInternalTidakBocor(t *testing.T) {
	svc := &fakeService{err: errors.New("pq: password authentication failed for user lentera")}
	_, body := doGET(t, newTestRouter(svc, 2), "/api/v1/cohorts/1/materials")

	if body.Error == nil || body.Error.Message != "internal server error" {
		t.Fatalf("pesan 500 harus generik, got %+v", body.Error)
	}
}

// ───────────────── test: input diteruskan ke service dengan benar ─────────────────

func TestMaterialHandler_List_MeneruskanQuery(t *testing.T) {
	svc := &fakeService{}
	w, _ := doGET(t, newTestRouter(svc, 2), "/api/v1/cohorts/7/materials?q=gin&status=submitted&late=true")

	if w.Code != 200 {
		t.Fatalf("want 200, got %d", w.Code)
	}
	if svc.gotUserID != 2 || svc.gotCohortID != 7 {
		t.Errorf("want user 2 cohort 7, got user %d cohort %d", svc.gotUserID, svc.gotCohortID)
	}
	if svc.gotQuery.Q != "gin" || svc.gotQuery.Status != "submitted" {
		t.Errorf("query tidak diteruskan: %+v", svc.gotQuery)
	}
	if svc.gotQuery.Late == nil || !*svc.gotQuery.Late {
		t.Errorf("late=true harus jadi pointer true, got %v", svc.gotQuery.Late)
	}
}

func TestMaterialHandler_List_TanpaFilterLate(t *testing.T) {
	svc := &fakeService{}
	doGET(t, newTestRouter(svc, 2), "/api/v1/cohorts/1/materials")

	if svc.gotQuery.Late != nil {
		t.Errorf("late tidak dikirim harus nil (tidak difilter), got %v", *svc.gotQuery.Late)
	}
}

// ───────────────── test: bentuk JSON sesuai kontrak ─────────────────

func TestMaterialHandler_List_BentukJSON(t *testing.T) {
	_, body := doGET(t, newTestRouter(&fakeService{}, 2), "/api/v1/cohorts/1/materials")

	var data []map[string]any
	if err := json.Unmarshal(body.Data, &data); err != nil {
		t.Fatalf("data harus array: %v", err)
	}
	if len(data) != 1 {
		t.Fatalf("want 1 item, got %d", len(data))
	}
	for _, key := range []string{"id", "title", "references", "my_progress", "position", "is_past_deadline"} {
		if _, ok := data[0][key]; !ok {
			t.Errorf("field %q tidak ada di response", key)
		}
	}
	if refs, _ := data[0]["references"].([]any); refs == nil {
		t.Errorf("references harus [] bukan null")
	}
}

// ───────────────── helper request dengan body ─────────────────

func doJSON(t *testing.T, r *gin.Engine, method, path, body string) (*httptest.ResponseRecorder, envelope) {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var env envelope
	if w.Body.Len() > 0 {
		if err := json.Unmarshal(w.Body.Bytes(), &env); err != nil {
			t.Fatalf("response bukan JSON valid: %v\nbody: %s", err, w.Body.String())
		}
	}
	return w, env
}

// ───────────────── CRUD: status code & validasi body ─────────────────

func TestMaterialHandler_CRUD_StatusCodes(t *testing.T) {
	const base = "/api/v1/cohorts/1/materials"

	tests := []struct {
		name          string
		method        string
		path          string
		body          string
		uid           uint
		svcErr        error
		wantStatus    int
		wantCode      string
		wantSvcCalled bool
	}{
		// create
		{name: "create sukses", method: "POST", path: base, body: `{"title":"Pengenalan Gin","guide":"baca docs"}`, uid: 1, wantStatus: 201, wantSvcCalled: true},
		{name: "create tanpa title", method: "POST", path: base, body: `{"guide":"baca docs"}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "create title terlalu pendek", method: "POST", path: base, body: `{"title":"ab","guide":"x"}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "create references bukan url", method: "POST", path: base, body: `{"title":"Gin","guide":"x","references":["bukan-url"]}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "create deadline format salah", method: "POST", path: base, body: `{"title":"Gin","guide":"x","deadline":"besok"}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "create body bukan JSON", method: "POST", path: base, body: `{title:`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "create oleh member", method: "POST", path: base, body: `{"title":"Gin","guide":"x"}`, uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},
		{name: "create tanpa login", method: "POST", path: base, body: `{"title":"Gin","guide":"x"}`, uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},

		// update
		{name: "update sukses", method: "PATCH", path: base + "/10", body: `{"title":"Gin Lanjutan"}`, uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "update materi tidak ada", method: "PATCH", path: base + "/999", body: `{"title":"Gin Lanjutan"}`, uid: 1, svcErr: response.ErrNotFound, wantStatus: 404, wantCode: "NOT_FOUND", wantSvcCalled: true},
		{name: "update materialId bukan angka", method: "PATCH", path: base + "/abc", body: `{"title":"Gin Lanjutan"}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},

		// delete
		{name: "delete sukses", method: "DELETE", path: base + "/10", uid: 1, wantStatus: 204, wantSvcCalled: true},
		{name: "delete oleh member", method: "DELETE", path: base + "/10", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},

		// reorder
		{name: "reorder sukses", method: "PUT", path: base + "/order", body: `{"material_ids":[3,1,2]}`, uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "reorder kosong", method: "PUT", path: base + "/order", body: `{"material_ids":[]}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "reorder id nol", method: "PUT", path: base + "/order", body: `{"material_ids":[0]}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "reorder ID tidak lengkap", method: "PUT", path: base + "/order", body: `{"material_ids":[1]}`, uid: 1, svcErr: response.ErrBadRequest, wantStatus: 400, wantCode: "BAD_INPUT", wantSvcCalled: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeService{err: tt.svcErr}
			w, body := doJSON(t, newTestRouter(svc, tt.uid), tt.method, tt.path, tt.body)

			if w.Code != tt.wantStatus {
				t.Fatalf("want status %d, got %d (body: %s)", tt.wantStatus, w.Code, w.Body.String())
			}
			if svc.called != tt.wantSvcCalled {
				t.Errorf("service called = %v, want %v", svc.called, tt.wantSvcCalled)
			}

			switch {
			case tt.wantStatus == http.StatusNoContent:
				if w.Body.Len() != 0 {
					t.Errorf("204 tidak boleh punya body, got %s", w.Body.String())
				}
			case tt.wantCode == "":
				if !body.Success {
					t.Errorf("want success envelope, got %s", w.Body.String())
				}
			default:
				if body.Error == nil || body.Error.Code != tt.wantCode {
					t.Errorf("want code %s, got %s", tt.wantCode, w.Body.String())
				}
			}
		})
	}
}

// ───────────────── PATCH: 3 kondisi deadline ─────────────────

func TestMaterialHandler_Update_Deadline(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantSet   bool
		wantValue bool // true = Value tidak nil
	}{
		{name: "deadline tidak dikirim", body: `{"title":"Gin Lanjutan"}`, wantSet: false, wantValue: false},
		{name: "deadline null (hapus)", body: `{"deadline":null}`, wantSet: true, wantValue: false},
		{name: "deadline diisi", body: `{"deadline":"2026-11-20T23:59:00Z"}`, wantSet: true, wantValue: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeService{}
			w, _ := doJSON(t, newTestRouter(svc, 1), http.MethodPatch, "/api/v1/cohorts/1/materials/10", tt.body)

			if w.Code != 200 {
				t.Fatalf("want 200, got %d (%s)", w.Code, w.Body.String())
			}
			got := svc.gotUpdate.Deadline
			if got.Set != tt.wantSet {
				t.Errorf("Set = %v, want %v", got.Set, tt.wantSet)
			}
			if (got.Value != nil) != tt.wantValue {
				t.Errorf("Value = %v, want non-nil=%v", got.Value, tt.wantValue)
			}
		})
	}
}

// ───────────────── POST: body diteruskan ke service ─────────────────

func TestMaterialHandler_Create_MeneruskanBody(t *testing.T) {
	svc := &fakeService{}
	doJSON(t, newTestRouter(svc, 1), http.MethodPost, "/api/v1/cohorts/5/materials",
		`{"title":"Gin","guide":"baca","references":["https://gin-gonic.com/docs/"],"deadline":"2026-11-20T23:59:00Z"}`)

	if svc.gotCohortID != 5 || svc.gotUserID != 1 {
		t.Errorf("want cohort 5 user 1, got cohort %d user %d", svc.gotCohortID, svc.gotUserID)
	}
	if svc.gotCreate.Title != "Gin" || len(svc.gotCreate.References) != 1 || svc.gotCreate.Deadline == nil {
		t.Errorf("body tidak diteruskan dengan benar: %+v", svc.gotCreate)
	}
}

func (f *fakeService) StartMaterial(_ context.Context, userID, cohortID, materialID uint) (*dto.ProgressResponse, error) {
	f.gotUserID, f.gotCohortID = userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.ProgressResponse{UserID: userID, MaterialID: materialID, Status: "in_progress"}, nil
}

func (f *fakeService) SubmitMaterial(_ context.Context, userID, cohortID, materialID uint, _ dto.SubmitMaterialRequest) (*dto.ProgressResponse, error) {
	f.gotUserID, f.gotCohortID = userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.ProgressResponse{UserID: userID, MaterialID: materialID, Status: "submitted"}, nil
}

func (f *fakeService) ReviewSubmission(_ context.Context, mentorID, cohortID, materialID, targetUserID uint, _ dto.ReviewSubmissionRequest) (*dto.ProgressResponse, error) {
	f.gotUserID, f.gotCohortID = mentorID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.ProgressResponse{UserID: targetUserID, MaterialID: materialID, Status: "completed"}, nil
}

func (f *fakeService) UploadPDF(_ context.Context, userID, cohortID uint, _, originalName string, data []byte) (*dto.UploadedFileResponse, error) {
	f.gotUserID, f.gotCohortID = userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.UploadedFileResponse{
		URL:         "https://example.supabase.co/storage/v1/object/public/lentera-files/" + originalName,
		FileName:    originalName,
		ContentType: "application/pdf",
		Size:        int64(len(data)),
	}, nil
}