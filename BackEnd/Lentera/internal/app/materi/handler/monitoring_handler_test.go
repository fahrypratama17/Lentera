package handler

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake service ─────────────────

type fakeMonitoringService struct {
	err error

	called      bool
	gotUserID   uint
	gotCohortID uint
	gotMemberID uint
	gotQuery    dto.MatrixQuery
}

func (f *fakeMonitoringService) Dashboard(_ context.Context, userID, cohortID uint) (*dto.DashboardResponse, error) {
	f.called, f.gotUserID, f.gotCohortID = true, userID, cohortID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.DashboardResponse{ActiveMemberCount: 2, MaterialCount: 3}, nil
}

func (f *fakeMonitoringService) Matrix(_ context.Context, userID, cohortID uint, q dto.MatrixQuery) (*dto.MatrixResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotQuery = true, userID, cohortID, q
	if f.err != nil {
		return nil, f.err
	}
	return &dto.MatrixResponse{
		Materials: []dto.MatrixMaterial{{ID: 10, Title: "Gin", Position: 1}},
		Rows: []dto.MatrixRow{{
			UserID: 2, Name: "Budi",
			Cells: []dto.MatrixCell{{MaterialID: 10, Status: "submitted"}},
		}},
	}, nil
}

func (f *fakeMonitoringService) MemberProgress(_ context.Context, userID, cohortID, memberID uint) (*dto.MemberProgressResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotMemberID = true, userID, cohortID, memberID
	if f.err != nil {
		return nil, f.err
	}
	return &dto.MemberProgressResponse{
		Member: dto.MemberBrief{UserID: memberID, Name: "Budi", Role: "member", Status: "active"},
		Items:  []dto.MemberProgressItem{},
	}, nil
}

// ───────────────── helper ─────────────────

func newMonitoringRouter(svc *fakeMonitoringService, uid uint) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	fakeAuth := func(c *gin.Context) {
		if uid != 0 {
			c.Set("uid", uid)
		}
		c.Next()
	}
	NewMonitoring(svc).RegisterRoutes(r.Group("/api/v1"), fakeAuth)
	return r
}

// ───────────────── status code & validasi ─────────────────

func TestMonitoringHandler_StatusCodes(t *testing.T) {
	const base = "/api/v1/cohorts/1"

	tests := []struct {
		name          string
		path          string
		uid           uint
		svcErr        error
		wantStatus    int
		wantCode      string
		wantSvcCalled bool
	}{
		// sukses
		{name: "dashboard sukses", path: base + "/dashboard", uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "matrix sukses", path: base + "/matrix", uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "matrix dengan filter", path: base + "/matrix?q=budi&status=submitted&material_id=10", uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "hasil member sukses", path: base + "/members/2/progress", uid: 1, wantStatus: 200, wantSvcCalled: true},

		// ditolak sebelum service
		{name: "dashboard tanpa login", path: base + "/dashboard", uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},
		{name: "matrix tanpa login", path: base + "/matrix", uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},
		{name: "cohortId bukan angka", path: "/api/v1/cohorts/abc/dashboard", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "userId nol", path: base + "/members/0/progress", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "status filter tidak valid", path: base + "/matrix?status=selesai", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "material_id bukan angka", path: base + "/matrix?material_id=abc", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},

		// ditolak oleh service
		{name: "member minta dashboard", path: base + "/dashboard", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},
		{name: "member minta matriks", path: base + "/matrix", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},
		{name: "member lihat hasil member lain", path: base + "/members/3/progress", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},
		{name: "material_id cohort lain", path: base + "/matrix?material_id=20", uid: 1, svcErr: response.ErrNotFound, wantStatus: 404, wantCode: "NOT_FOUND", wantSvcCalled: true},
		{name: "member tidak ditemukan", path: base + "/members/99/progress", uid: 1, svcErr: response.ErrNotFound, wantStatus: 404, wantCode: "NOT_FOUND", wantSvcCalled: true},
		{name: "error tak terduga", path: base + "/dashboard", uid: 1, svcErr: errors.New("pq: connection refused"), wantStatus: 500, wantCode: "INTERNAL", wantSvcCalled: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeMonitoringService{err: tt.svcErr}
			w, body := doGET(t, newMonitoringRouter(svc, tt.uid), tt.path)

			if w.Code != tt.wantStatus {
				t.Fatalf("want status %d, got %d (body: %s)", tt.wantStatus, w.Code, w.Body.String())
			}
			if svc.called != tt.wantSvcCalled {
				t.Errorf("service called = %v, want %v", svc.called, tt.wantSvcCalled)
			}

			if tt.wantCode == "" {
				if !body.Success || body.Error != nil {
					t.Errorf("want success envelope, got %s", w.Body.String())
				}
				return
			}
			if body.Error == nil || body.Error.Code != tt.wantCode {
				t.Errorf("want code %s, got %s", tt.wantCode, w.Body.String())
			}
		})
	}
}

// ───────────────── input diteruskan ke service ─────────────────

func TestMonitoringHandler_Matrix_MeneruskanQuery(t *testing.T) {
	svc := &fakeMonitoringService{}
	doGET(t, newMonitoringRouter(svc, 1), "/api/v1/cohorts/7/matrix?q=budi&status=submitted&material_id=10")

	if svc.gotUserID != 1 || svc.gotCohortID != 7 {
		t.Errorf("want user 1 cohort 7, got user %d cohort %d", svc.gotUserID, svc.gotCohortID)
	}
	if svc.gotQuery.Q != "budi" || svc.gotQuery.Status != "submitted" || svc.gotQuery.MaterialID != 10 {
		t.Errorf("query tidak diteruskan: %+v", svc.gotQuery)
	}
}

func TestMonitoringHandler_MemberProgress_MeneruskanID(t *testing.T) {
	svc := &fakeMonitoringService{}
	doGET(t, newMonitoringRouter(svc, 1), "/api/v1/cohorts/7/members/42/progress")

	if svc.gotUserID != 1 || svc.gotCohortID != 7 || svc.gotMemberID != 42 {
		t.Errorf("want user 1 cohort 7 member 42, got user %d cohort %d member %d",
			svc.gotUserID, svc.gotCohortID, svc.gotMemberID)
	}
}

// ───────────────── bentuk JSON sesuai kontrak ─────────────────

func TestMonitoringHandler_Matrix_BentukJSON(t *testing.T) {
	_, body := doGET(t, newMonitoringRouter(&fakeMonitoringService{}, 1), "/api/v1/cohorts/1/matrix")

	var data struct {
		Materials []map[string]any `json:"materials"`
		Rows      []struct {
			UserID         uint             `json:"user_id"`
			Name           string           `json:"name"`
			CompletionRate *float64         `json:"completion_rate"`
			Cells          []map[string]any `json:"cells"`
		} `json:"rows"`
	}
	if err := json.Unmarshal(body.Data, &data); err != nil {
		t.Fatalf("data bukan bentuk matriks: %v", err)
	}
	if len(data.Materials) != 1 || len(data.Rows) != 1 {
		t.Fatalf("want 1 materi 1 baris, got %d %d", len(data.Materials), len(data.Rows))
	}
	if data.Rows[0].CompletionRate == nil {
		t.Error("field completion_rate tidak ada di baris")
	}
	for _, key := range []string{"material_id", "progress_id", "status", "is_late"} {
		if _, ok := data.Rows[0].Cells[0][key]; !ok {
			t.Errorf("field %q tidak ada di sel matriks", key)
		}
	}
}

func TestMonitoringHandler_Dashboard_BentukJSON(t *testing.T) {
	_, body := doGET(t, newMonitoringRouter(&fakeMonitoringService{}, 1), "/api/v1/cohorts/1/dashboard")

	var data map[string]any
	if err := json.Unmarshal(body.Data, &data); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"active_member_count", "material_count", "pending_review_count", "completed_count", "completion_rate", "late_count"} {
		if _, ok := data[key]; !ok {
			t.Errorf("field %q tidak ada di dashboard", key)
		}
	}
}
