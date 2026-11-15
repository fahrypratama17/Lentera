package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORS mengizinkan origin yang ada di daftar saja.
// Origin dikembalikan persis (bukan "*") supaya bisa dipakai bareng Authorization/credentials.
func CORS(origins []string) gin.HandlerFunc {
	allowed := make(map[string]struct{}, len(origins))
	for _, o := range origins {
		o = strings.TrimRight(strings.TrimSpace(o), "/")
		if o != "" {
			allowed[o] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if _, ok := allowed[origin]; ok {
			h := c.Writer.Header()
			h.Set("Access-Control-Allow-Origin", origin)
			h.Add("Vary", "Origin")
			h.Set("Access-Control-Allow-Credentials", "true")
			h.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
			h.Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-Request-ID")
			h.Set("Access-Control-Expose-Headers", "X-Request-ID")
			h.Set("Access-Control-Max-Age", "600")
		}

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}