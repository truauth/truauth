package middleware

import (
	"github.com/gin-gonic/gin"
)

// EnableCORS middlewear to enable cors
func EnableCORS(ctx *gin.Context) {
	(ctx.Writer).Header().Set("Access-Control-Allow-Origin", "*")
}
