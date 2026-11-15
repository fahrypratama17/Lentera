package repository

import (
	"errors"
	"testing"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

func TestSameIDSet(t *testing.T) {
	existing := []uint{10, 11, 12}

	tests := []struct {
		name    string
		ordered []uint
		wantErr bool
	}{
		{name: "lengkap urutan baru", ordered: []uint{12, 10, 11}, wantErr: false},
		{name: "kurang satu", ordered: []uint{12, 10}, wantErr: true},
		{name: "ada ID cohort lain", ordered: []uint{12, 10, 99}, wantErr: true},
		{name: "duplikat", ordered: []uint{12, 12, 10}, wantErr: true},
		{name: "kelebihan", ordered: []uint{10, 11, 12, 13}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := sameIDSet(existing, tt.ordered)
			if tt.wantErr {
				if !errors.Is(err, response.ErrBadRequest) {
					t.Fatalf("want ErrBadRequest, got %v", err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
		})
	}
}
