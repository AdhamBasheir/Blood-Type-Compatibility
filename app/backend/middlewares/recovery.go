package middleware

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if rec := recover(); rec != nil {
				// Capture stack trace
				stack := string(debug.Stack())

				// Log the panic with request details
				fields := getBaseRequestLogFields(ctx)
				fields["recovered"] = fmt.Sprintf("%v", rec)
				fields["stack"] = string(bytes.TrimSpace([]byte(stack)))

				logrus.WithFields(fields).Error("panic recovered")

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
				})
			}
		}()

		ctx.Next()
	}
}
