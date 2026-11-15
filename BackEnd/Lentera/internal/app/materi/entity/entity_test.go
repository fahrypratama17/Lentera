package entity

import (
	"errors"
	"testing"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

var now = time.Date(2026, 11, 15, 12, 0, 0, 0, time.UTC)

func TestProgress_Transitions(t *testing.T) {
	note := "bagus"

	tests := []struct {
		name       string
		from       string
		action     func(p *Progress) error
		wantErr    error
		wantStatus string
	}{
		// berhasil
		{name: "approve dari submitted", from: StatusSubmitted, action: func(p *Progress) error { return p.Approve(1, &note, now) }, wantStatus: StatusCompleted},
		{name: "approve tanpa catatan", from: StatusSubmitted, action: func(p *Progress) error { return p.Approve(1, nil, now) }, wantStatus: StatusCompleted},
		{name: "revisi dari submitted", from: StatusSubmitted, action: func(p *Progress) error { return p.RequestRevision(1, "tambah test", now) }, wantStatus: StatusInProgress},
		{name: "reopen dari completed", from: StatusCompleted, action: func(p *Progress) error { return p.Reopen(1, "ada requirement baru", now) }, wantStatus: StatusInProgress},

		// ditolak — status harus tetap
		{name: "approve dari in_progress", from: StatusInProgress, action: func(p *Progress) error { return p.Approve(1, nil, now) }, wantErr: response.ErrConflict, wantStatus: StatusInProgress},
		{name: "approve dari not_started", from: StatusNotStarted, action: func(p *Progress) error { return p.Approve(1, nil, now) }, wantErr: response.ErrConflict, wantStatus: StatusNotStarted},
		{name: "approve dua kali", from: StatusCompleted, action: func(p *Progress) error { return p.Approve(1, nil, now) }, wantErr: response.ErrConflict, wantStatus: StatusCompleted},
		{name: "revisi dari completed", from: StatusCompleted, action: func(p *Progress) error { return p.RequestRevision(1, "x", now) }, wantErr: response.ErrConflict, wantStatus: StatusCompleted},
		{name: "revisi dari in_progress", from: StatusInProgress, action: func(p *Progress) error { return p.RequestRevision(1, "x", now) }, wantErr: response.ErrConflict, wantStatus: StatusInProgress},
		{name: "reopen dari submitted", from: StatusSubmitted, action: func(p *Progress) error { return p.Reopen(1, "x", now) }, wantErr: response.ErrConflict, wantStatus: StatusSubmitted},
		{name: "revisi tanpa catatan", from: StatusSubmitted, action: func(p *Progress) error { return p.RequestRevision(1, "   ", now) }, wantErr: response.ErrBadRequest, wantStatus: StatusSubmitted},
		{name: "reopen tanpa catatan", from: StatusCompleted, action: func(p *Progress) error { return p.Reopen(1, "", now) }, wantErr: response.ErrBadRequest, wantStatus: StatusCompleted},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Progress{Status: tt.from}
			err := tt.action(p)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("want %v, got %v", tt.wantErr, err)
				}
				if p.ReviewDecision != nil {
					t.Error("kalau ditolak, data review tidak boleh terisi")
				}
			} else if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}

			if p.Status != tt.wantStatus {
				t.Errorf("status: want %s, got %s", tt.wantStatus, p.Status)
			}
		})
	}
}

func TestProgress_ApproveIsiCompletedAt(t *testing.T) {
	p := &Progress{Status: StatusSubmitted}
	if err := p.Approve(1, nil, now); err != nil {
		t.Fatal(err)
	}
	if p.CompletedAt == nil || !p.CompletedAt.Equal(now) {
		t.Errorf("completed_at harus terisi waktu approve, got %v", p.CompletedAt)
	}
	if p.ReviewNote != nil {
		t.Error("approve tanpa catatan: review_note harus nil")
	}
}

func TestProgress_ReviewTercatat(t *testing.T) {
	p := &Progress{Status: StatusSubmitted}
	if err := p.RequestRevision(7, "  tambah error handling  ", now); err != nil {
		t.Fatal(err)
	}
	if p.ReviewDecision == nil || *p.ReviewDecision != DecisionRevisionRequested {
		t.Errorf("decision salah: %v", p.ReviewDecision)
	}
	if p.ReviewNote == nil || *p.ReviewNote != "tambah error handling" {
		t.Errorf("catatan harus di-trim, got %v", p.ReviewNote)
	}
	if p.ReviewerID == nil || *p.ReviewerID != 7 || p.ReviewedAt == nil || !p.ReviewedAt.Equal(now) {
		t.Error("reviewer & waktu review harus tercatat")
	}
}

func TestProgress_ReopenKosongkanCompletedAt(t *testing.T) {
	done := now.Add(-time.Hour)
	p := &Progress{Status: StatusCompleted, CompletedAt: &done}
	if err := p.Reopen(1, "ulang", now); err != nil {
		t.Fatal(err)
	}
	if p.CompletedAt != nil {
		t.Error("completed_at harus dikosongkan setelah reopen")
	}
	if *p.ReviewDecision != DecisionReopened {
		t.Errorf("decision harus reopened, got %s", *p.ReviewDecision)
	}
}

func TestProgress_ReviewBaruMenimpaYangLama(t *testing.T) {
	p := &Progress{Status: StatusSubmitted}
	_ = p.RequestRevision(1, "revisi pertama", now)

	// member submit ulang
	p.Status = StatusSubmitted
	later := now.Add(time.Hour)
	if err := p.Approve(2, nil, later); err != nil {
		t.Fatal(err)
	}
	if *p.ReviewDecision != DecisionApproved || *p.ReviewerID != 2 || !p.ReviewedAt.Equal(later) {
		t.Error("keputusan terakhir harus menimpa keputusan sebelumnya")
	}
}
