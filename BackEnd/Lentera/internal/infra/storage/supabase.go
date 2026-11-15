package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SupabaseStorage struct {
	baseURL    string
	serviceKey string
	bucket     string
	httpClient *http.Client
}

func NewSupabaseStorage(baseURL, serviceKey, bucket string) *SupabaseStorage {
	return &SupabaseStorage{
		baseURL:    strings.TrimRight(baseURL, "/"),
		serviceKey: serviceKey,
		bucket:     bucket,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

// UploadPDF memenuhi interface contract.FileUploader untuk mengunggah file PDF ke Supabase Storage.
func (s *SupabaseStorage) UploadPDF(ctx context.Context, objectPath string, data []byte) (string, error) {
	return s.Upload(ctx, objectPath, "application/pdf", data)
}

func (s *SupabaseStorage) Upload(ctx context.Context, objectPath, contentType string, data []byte) (string, error) {
	if s.baseURL == "" || s.serviceKey == "" {
		return "", fmt.Errorf("supabase storage belum dikonfigurasi")
	}

	encodedBucket := url.PathEscape(s.bucket)
	encodedPath := escapeObjectPath(objectPath)
	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", s.baseURL, encodedBucket, encodedPath)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uploadURL, bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("apikey", s.serviceKey)
	req.Header.Set("Authorization", "Bearer "+s.serviceKey)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("x-upsert", "true")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("supabase upload gagal (%d): %s", resp.StatusCode, string(body))
	}

	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", s.baseURL, encodedBucket, encodedPath)
	return publicURL, nil
}

func escapeObjectPath(objectPath string) string {
	clean := strings.TrimLeft(objectPath, "/")
	parts := strings.Split(clean, "/")
	for i, part := range parts {
		parts[i] = url.PathEscape(part)
	}
	return strings.Join(parts, "/")
}