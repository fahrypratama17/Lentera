package handler

import (
	"net/http"

	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/contract"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/internal/app/user/dto"
	"github.com/fahrypratama17/Lentera/BackEnd/Lentera/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc contract.UserService
}

func New(svc contract.UserService) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.svc.Register(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusCreated, res)
}

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	res, err := h.svc.RefreshToken(c.Request.Context(), req)
	if err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	var req dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err)
		return
	}
	if err := h.svc.Logout(c.Request.Context(), req); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, http.StatusOK, gin.H{"message": "logged out"})
}