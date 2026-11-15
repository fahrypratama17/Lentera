package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// ───────────────── fake service ─────────────────

type fakeReviewService struct {
	err error

	called        bool
	gotUserID     uint
	gotCohortID   uint
	gotProgressID uint
	gotQuery      dto.ListReviewsQuery
	gotApprove    dto.ApproveRequest
	gotNote       dto.ReviewNoteRequest
}

func (f *fakeReviewService) item(progressID uint) *dto.ReviewItemResponse {
	id := progressID
	return &dto.ReviewItemResponse{
		ProgressResponse: dto.ProgressResponse{ID: &id, Status: "completed"},
		Member:           dto.ReviewMember{UserID: 2, Name: "Budi", Status: "active"},
		Material:         dto.MatrixMaterial{ID: 10, Title: "Gin"},
	}
}

func (f *fakeReviewService) ListReviews(_ context.Context, userID, cohortID uint, q dto.ListReviewsQuery) ([]dto.ReviewItemResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotQuery = true, userID, cohortID, q
	if f.err != nil {
		return nil, f.err
	}
	return []dto.ReviewItemResponse{*f.item(1)}, nil
}

func (f *fakeReviewService) GetReviewItem(_ context.Context, userID, cohortID, progressID uint) (*dto.ReviewItemResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotProgressID = true, userID, cohortID, progressID
	if f.err != nil {
		return nil, f.err
	}
	return f.item(progressID), nil
}

func (f *fakeReviewService) Approve(_ context.Context, userID, cohortID, progressID uint, req dto.ApproveRequest) (*dto.ReviewItemResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotProgressID, f.gotApprove = true, userID, cohortID, progressID, req
	if f.err != nil {
		return nil, f.err
	}
	return f.item(progressID), nil
}

func (f *fakeReviewService) RequestRevision(_ context.Context, userID, cohortID, progressID uint, req dto.ReviewNoteRequest) (*dto.ReviewItemResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotProgressID, f.gotNote = true, userID, cohortID, progressID, req
	if f.err != nil {
		return nil, f.err
	}
	return f.item(progressID), nil
}

func (f *fakeReviewService) Reopen(_ context.Context, userID, cohortID, progressID uint, req dto.ReviewNoteRequest) (*dto.ReviewItemResponse, error) {
	f.called, f.gotUserID, f.gotCohortID, f.gotProgressID, f.gotNote = true, userID, cohortID, progressID, req
	if f.err != nil {
		return nil, f.err
	}
	return f.item(progressID), nil
}

// ───────────────── helper ─────────────────

func newReviewRouter(svc *fakeReviewService, uid uint) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()

	fakeAuth := func(c *gin.Context) {
		if uid != 0 {
			c.Set("uid", uid)
		}
		c.Next()
	}
	NewReview(svc).RegisterRoutes(r.Group("/api/v1"), fakeAuth)
	return r
}

// ───────────────── status code & validasi ─────────────────

func TestReviewHandler_StatusCodes(t *testing.T) {
	const base = "/api/v1/cohorts/1"
	conflict := fmt.Errorf("%w: progres sudah direview", response.ErrConflict)

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
		// daftar & detail
		{name: "list sukses", method: "GET", path: base + "/reviews", uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "list tanpa login", method: "GET", path: base + "/reviews", uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},
		{name: "list status tidak valid", method: "GET", path: base + "/reviews?status=not_started", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "list material_id bukan angka", method: "GET", path: base + "/reviews?material_id=abc", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "list oleh member", method: "GET", path: base + "/reviews", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},
		{name: "detail sukses", method: "GET", path: base + "/progress/7", uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "detail progressId nol", method: "GET", path: base + "/progress/0", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "detail tidak ditemukan", method: "GET", path: base + "/progress/999", uid: 1, svcErr: response.ErrNotFound, wantStatus: 404, wantCode: "NOT_FOUND", wantSvcCalled: true},
		{name: "detail milik member lain", method: "GET", path: base + "/progress/7", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},

		// approve
		{name: "approve dengan catatan", method: "POST", path: base + "/progress/1/approve", body: `{"note":"mantap"}`, uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "approve tanpa body", method: "POST", path: base + "/progress/1/approve", body: "", uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "approve body rusak", method: "POST", path: base + "/progress/1/approve", body: `{note:`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "approve oleh member", method: "POST", path: base + "/progress/1/approve", body: "", uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN", wantSvcCalled: true},
		{name: "approve sudah direview (race)", method: "POST", path: base + "/progress/1/approve", body: "", uid: 1, svcErr: conflict, wantStatus: 409, wantCode: "CONFLICT", wantSvcCalled: true},
		{name: "approve tanpa login", method: "POST", path: base + "/progress/1/approve", body: "", uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},

		// request revision
		{name: "revisi sukses", method: "POST", path: base + "/progress/1/request-revision", body: `{"note":"tambah test"}`, uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "revisi tanpa note", method: "POST", path: base + "/progress/1/request-revision", body: `{}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "revisi note kosong", method: "POST", path: base + "/progress/1/request-revision", body: `{"note":""}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "revisi transisi tidak valid", method: "POST", path: base + "/progress/3/request-revision", body: `{"note":"x"}`, uid: 1, svcErr: conflict, wantStatus: 409, wantCode: "CONFLICT", wantSvcCalled: true},

		// reopen
		{name: "reopen sukses", method: "POST", path: base + "/progress/2/reopen", body: `{"note":"ada requirement baru"}`, uid: 1, wantStatus: 200, wantSvcCalled: true},
		{name: "reopen tanpa body", method: "POST", path: base + "/progress/2/reopen", body: "", uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeReviewService{err: tt.svcErr}
			w, body := doJSON(t, newReviewRouter(svc, tt.uid), tt.method, tt.path, tt.body)

			if w.Code != tt.wantStatus {
				t.Fatalf("want status %d, got %d (body: %s)", tt.wantStatus, w.Code, w.Body.String())
			}
			if svc.called != tt.wantSvcCalled {
				t.Errorf("service called = %v, want %v", svc.called, tt.wantSvcCalled)
			}

			if tt.wantCode == "" {
				if !body.Success {
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

func TestReviewHandler_List_MeneruskanQuery(t *testing.T) {
	svc := &fakeReviewService{}
	doGET(t, newReviewRouter(svc, 1), "/api/v1/cohorts/5/reviews?status=completed&material_id=10&q=budi&include_inactive=true")

	if svc.gotUserID != 1 || svc.gotCohortID != 5 {
		t.Errorf("want user 1 cohort 5, got user %d cohort %d", svc.gotUserID, svc.gotCohortID)
	}
	q := svc.gotQuery
	if q.Status != "completed" || q.MaterialID != 10 || q.Q != "budi" || !q.IncludeInactive {
		t.Errorf("query tidak diteruskan: %+v", q)
	}
}

func TestReviewHandler_Approve_MeneruskanCatatan(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		wantNote *string
	}{
		{name: "dengan catatan", body: `{"note":"mantap"}`, wantNote: strPtrH("mantap")},
		{name: "tanpa body", body: "", wantNote: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeReviewService{}
			doJSON(t, newReviewRouter(svc, 1), http.MethodPost, "/api/v1/cohorts/5/progress/42/approve", tt.body)

			if svc.gotCohortID != 5 || svc.gotProgressID != 42 {
				t.Errorf("want cohort 5 progres 42, got %d %d", svc.gotCohortID, svc.gotProgressID)
			}
			got := svc.gotApprove.Note
			if (got == nil) != (tt.wantNote == nil) || (got != nil && *got != *tt.wantNote) {
				t.Errorf("note: want %v, got %v", tt.wantNote, got)
			}
		})
	}
}

func TestReviewHandler_RequestRevision_MeneruskanCatatan(t *testing.T) {
	svc := &fakeReviewService{}
	doJSON(t, newReviewRouter(svc, 1), http.MethodPost, "/api/v1/cohorts/1/progress/7/request-revision", `{"note":"tambah test"}`)

	if svc.gotProgressID != 7 || svc.gotNote.Note != "tambah test" {
		t.Errorf("want progres 7 note 'tambah test', got %d %q", svc.gotProgressID, svc.gotNote.Note)
	}
}

func TestReviewHandler_ErrorInternalTidakBocor(t *testing.T) {
	svc := &fakeReviewService{err: errors.New("pq: deadlock detected")}
	_, body := doJSON(t, newReviewRouter(svc, 1), http.MethodPost, "/api/v1/cohorts/1/progress/1/approve", "")

	if body.Error == nil || body.Error.Message != "internal server error" {
		t.Fatalf("pesan 500 harus generik, got %+v", body.Error)
	}
}

func strPtrH(s string) *string { return &s }
