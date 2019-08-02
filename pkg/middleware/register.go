package middleware

import (
	"github.com/gin-gonic/gin"
)

// Middlewear middlewear type handler
type Middlewear func(ctx *gin.Context)

// Register runs the desiered middlewear with the handler.
func Register(handler gin.HandlerFunc, middlewears ...Middlewear) gin.HandlerFunc {
	for _, middlewear := range middlewears {
		handler = func(handler gin.HandlerFunc) gin.HandlerFunc {
			return gin.HandlerFunc(func(ctx *gin.Context) {
				middlewear(ctx)

				handler(ctx)
			})
		}(handler)
	}

	return handler
}
