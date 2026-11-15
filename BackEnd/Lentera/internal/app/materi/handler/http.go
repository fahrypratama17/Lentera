package handler

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"
)

func writeOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func writeCreated(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": data})
}

func writeNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

func writeError(c *gin.Context, err error) {
	status, code := http.StatusInternalServerError, "INTERNAL"

	switch {
	case errors.Is(err, response.ErrBadRequest):
		status, code = http.StatusBadRequest, "BAD_INPUT"
	case errors.Is(err, response.ErrUnauthorized):
		status, code = http.StatusUnauthorized, "UNAUTHORIZED"
	case errors.Is(err, response.ErrForbidden):
		status, code = http.StatusForbidden, "FORBIDDEN"
	case errors.Is(err, response.ErrNotFound):
		status, code = http.StatusNotFound, "NOT_FOUND"
	case errors.Is(err, response.ErrConflict):
		status, code = http.StatusConflict, "CONFLICT"
	case errors.Is(err, response.ErrNotImplemented):
		status, code = http.StatusNotImplemented, "NOT_IMPLEMENTED"
	}

	msg := err.Error()
	if status == http.StatusInternalServerError {
		slog.ErrorContext(c.Request.Context(), "internal error", "error", err, "route", c.FullPath())
		msg = "internal server error" // jangan bocorkan detail DB ke client
	}

	c.AbortWithStatusJSON(status, gin.H{
		"success": false,
		"error":   gin.H{"code": code, "message": msg},
	})
}

// ⚠️ SESUAIKAN: key & tipe harus sama dengan yang di-set middleware/jwt.go
func currentUser(c *gin.Context) (uint, bool) {
	id := c.GetUint("uid")
	if id == 0 {
		writeError(c, response.ErrUnauthorized)
		return 0, false
	}
	return id, true
}

func paramID(c *gin.Context, name string) (uint, error) {
	v, err := strconv.ParseUint(c.Param(name), 10, 64)
	if err != nil || v == 0 {
		return 0, fmt.Errorf("%w: %s harus angka positif", response.ErrBadRequest, name)
	}
	return uint(v), nil
}
