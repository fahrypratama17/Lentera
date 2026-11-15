package handler

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(api *gin.RouterGroup, auth gin.HandlerFunc) {
	a := api.Group("/auth")
	a.POST("/register", h.Register)
	a.POST("/login", h.Login)
	a.POST("/refresh", h.RefreshToken)
	a.POST("/logout", h.Logout)

	api.GET("/me", auth, h.Me)
}
