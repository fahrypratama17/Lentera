package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake service ─────────────────

type fakeService struct {
	err error

	// dicatat supaya bisa dicek handler meneruskan input dengan benar
	gotUserID   uint
	gotCohortID uint
	gotQuery    dto.ListMaterialsQuery
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

// ───────────────── helper ─────────────────

// uid = 0 artinya simulasi request tanpa login
func newTestRouter(svc *fakeService, uid uint) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	// pengganti middleware JWT: set uid seperti yang dilakukan middleware asli
	protected := r.Group("/api/v1", func(c *gin.Context) {
		if uid != 0 {
			c.Set("uid", uid)
		}
		c.Next()
	})
	NewMaterialHandler(svc).RegisterRoutes(protected)
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
