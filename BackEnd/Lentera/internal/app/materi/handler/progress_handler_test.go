package handler

import (
	"net/http"
	"testing"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

func TestMaterialHandler_Progress_StatusCodes(t *testing.T) {
	const base = "/api/v1/cohorts/1/materials/10"

	tests := []struct {
		name       string
		path       string
		body       string
		uid        uint
		svcErr     error
		wantStatus int
		wantCode   string
	}{
		// start
		{name: "start sukses", path: base + "/start", uid: 2, wantStatus: 200},
		{name: "start sudah submitted", path: base + "/start", uid: 2, svcErr: response.ErrConflict, wantStatus: 409, wantCode: "CONFLICT"},
		{name: "start oleh mentor", path: base + "/start", uid: 1, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN"},
		{name: "start tanpa login", path: base + "/start", uid: 0, wantStatus: 401, wantCode: "UNAUTHORIZED"},

		// submit
		{name: "submit sukses", path: base + "/submit", body: `{"content":"selesai","work_url":"https://github.com/lentera/repo"}`, uid: 2, wantStatus: 200},
		{name: "submit work_url bukan url", path: base + "/submit", body: `{"work_url":"bukan-url"}`, uid: 2, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "submit body invalid json", path: base + "/submit", body: `{content:`, uid: 2, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "submit belum start", path: base + "/submit", body: `{"content":"selesai"}`, uid: 2, svcErr: response.ErrConflict, wantStatus: 409, wantCode: "CONFLICT"},

		// review
		{name: "review approve sukses", path: base + "/submissions/2/review", body: `{"decision":"approved"}`, uid: 1, wantStatus: 200},
		{name: "review revisi sukses", path: base + "/submissions/2/review", body: `{"decision":"revision_requested","note":"perbaiki"}`, uid: 1, wantStatus: 200},
		{name: "review decision tidak valid", path: base + "/submissions/2/review", body: `{"decision":"rejected"}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "review userId nol", path: base + "/submissions/0/review", body: `{"decision":"approved"}`, uid: 1, wantStatus: 400, wantCode: "BAD_INPUT"},
		{name: "review oleh member", path: base + "/submissions/2/review", body: `{"decision":"approved"}`, uid: 2, svcErr: response.ErrForbidden, wantStatus: 403, wantCode: "FORBIDDEN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &fakeService{err: tt.svcErr}
			w, body := doJSON(t, newTestRouter(svc, tt.uid), http.MethodPost, tt.path, tt.body)

			if w.Code != tt.wantStatus {
				t.Fatalf("want status %d, got %d (body: %s)", tt.wantStatus, w.Code, w.Body.String())
			}
			if tt.wantCode == "" {
				if !body.Success {
					t.Fatalf("want success, got %s", w.Body.String())
				}
				return
			}
			if body.Error == nil || body.Error.Code != tt.wantCode {
				t.Errorf("want code %s, got %s", tt.wantCode, w.Body.String())
			}
		})
	}
}