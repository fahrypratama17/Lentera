package handler

import (
	"fmt"
	"io"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/materi/service"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

// POST /cohorts/:id/materials/upload
// Form-Data:
//   - purpose: "material" (Mentor) | "submission" (Member)
//   - file:    <file .pdf maksimal 10MB>
func (h *MaterialHandler) UploadPDF(c *gin.Context) {
	userID, ok := currentUser(c)
	if !ok {
		return
	}
	cohortID, err := paramID(c, "id")
	if err != nil {
		writeError(c, err)
		return
	}

	purpose := c.PostForm("purpose")
	fileHeader, err := c.FormFile("file")
	if err != nil {
		writeError(c, fmt.Errorf("%w: field 'file' wajib diisi", response.ErrBadRequest))
		return
	}
	if fileHeader.Size > service.MaxPDFSize {
		writeError(c, fmt.Errorf("%w: ukuran file maksimal 10 MB", response.ErrBadRequest))
		return
	}

	f, err := fileHeader.Open()
	if err != nil {
		writeError(c, err)
		return
	}
	defer f.Close()

	data, err := io.ReadAll(io.LimitReader(f, service.MaxPDFSize+1))
	if err != nil {
		writeError(c, err)
		return
	}

	res, err := h.svc.UploadPDF(c.Request.Context(), userID, cohortID, purpose, fileHeader.Filename, data)
	if err != nil {
		writeError(c, err)
		return
	}
	writeCreated(c, res)
}