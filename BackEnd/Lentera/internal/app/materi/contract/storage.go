package contract

import "context"

type FileUploader interface {
	UploadPDF(ctx context.Context, objectPath string, data []byte) (string, error)
}