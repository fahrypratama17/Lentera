package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/entity"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

const MaxPDFSize = 10 << 20 // 10 MB

func (s *materialService) UploadPDF(ctx context.Context, userID, cohortID uint, purpose, originalName string, data []byte) (*dto.UploadedFileResponse, error) {
	if s.uploader == nil {
		return nil, response.ErrNotImplemented
	}

	role, err := s.repo.GetActiveRole(ctx, cohortID, userID)
	if err != nil {
		return nil, err
	}

	// Pisahkan folder penyimpanan & hak akses berdasarkan tujuan upload
	switch purpose {
	case "material":
		if role != entity.RoleMentor {
			return nil, fmt.Errorf("%w: hanya mentor yang dapat mengunggah file materi", response.ErrForbidden)
		}
	case "submission":
		if role != entity.RoleMember {
			return nil, fmt.Errorf("%w: hanya member yang dapat mengunggah file pengerjaan", response.ErrForbidden)
		}
	default:
		return nil, fmt.Errorf("%w: purpose harus 'material' atau 'submission'", response.ErrBadRequest)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("%w: file tidak boleh kosong", response.ErrBadRequest)
	}
	if len(data) > MaxPDFSize {
		return nil, fmt.Errorf("%w: ukuran file maksimal 10 MB", response.ErrBadRequest)
	}

	// 1. Cek ekstensi file
	if !strings.EqualFold(filepath.Ext(originalName), ".pdf") {
		return nil, fmt.Errorf("%w: format file harus PDF (.pdf)", response.ErrBadRequest)
	}

	// 2. Cek magic bytes (header %PDF-) agar file non-PDF yang di-rename tetap ditolak
	mimeType := http.DetectContentType(data)
	if mimeType != "application/pdf" {
		return nil, fmt.Errorf("%w: isi file bukan dokumen PDF yang valid", response.ErrBadRequest)
	}

	objectPath := fmt.Sprintf("cohorts/%d/%s/%d-%s.pdf", cohortID, purpose, userID, randomHex(12))
	publicURL, err := s.uploader.UploadPDF(ctx, objectPath, data)
	if err != nil {
		return nil, err
	}

	return &dto.UploadedFileResponse{
		URL:         publicURL,
		FileName:    filepath.Base(originalName),
		ContentType: "application/pdf",
		Size:        int64(len(data)),
	}, nil
}

func randomHex(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}