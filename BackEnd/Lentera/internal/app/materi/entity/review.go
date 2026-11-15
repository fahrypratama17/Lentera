package entity

import (
	"fmt"
	"strings"
	"time"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

const (
	DecisionApproved          = "approved"
	DecisionRevisionRequested = "revision_requested"
	DecisionReopened          = "reopened"
)

var (
	ErrInvalidTransition = fmt.Errorf("%w: transisi status tidak valid", response.ErrConflict)
	ErrNoteRequired      = fmt.Errorf("%w: catatan wajib diisi", response.ErrBadRequest)
)

// Approve: submitted → completed (catatan opsional)
func (p *Progress) Approve(reviewerID uint, note *string, now time.Time) error {
	if p.Status != StatusSubmitted {
		return ErrInvalidTransition
	}
	t := now
	p.Status = StatusCompleted
	p.CompletedAt = &t
	p.setReview(DecisionApproved, note, reviewerID, now)
	return nil
}

// RequestRevision: submitted → in_progress (catatan wajib)
func (p *Progress) RequestRevision(reviewerID uint, note string, now time.Time) error {
	if p.Status != StatusSubmitted {
		return ErrInvalidTransition
	}
	n := strings.TrimSpace(note)
	if n == "" {
		return ErrNoteRequired
	}
	p.Status = StatusInProgress
	p.setReview(DecisionRevisionRequested, &n, reviewerID, now)
	return nil
}

// Reopen: completed → in_progress (catatan wajib)
func (p *Progress) Reopen(reviewerID uint, note string, now time.Time) error {
	if p.Status != StatusCompleted {
		return ErrInvalidTransition
	}
	n := strings.TrimSpace(note)
	if n == "" {
		return ErrNoteRequired
	}
	p.Status = StatusInProgress
	p.CompletedAt = nil
	p.setReview(DecisionReopened, &n, reviewerID, now)
	return nil
}

func (p *Progress) setReview(decision string, note *string, reviewerID uint, now time.Time) {
	d, r, t := decision, reviewerID, now
	p.ReviewDecision = &d
	p.ReviewNote = note
	p.ReviewerID = &r
	p.ReviewedAt = &t
}
