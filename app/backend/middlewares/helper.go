package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func getBaseRequestLogFields(ctx *gin.Context) logrus.Fields {
	return logrus.Fields{
		"status":      ctx.Writer.Status(),
		"method":      ctx.Request.Method,
		"path":        ctx.Request.URL.Path,
		"ip":          ctx.ClientIP(),
		"userAgent":   ctx.Request.UserAgent(),
		"contentType": ctx.ContentType(),
	}
}

func logRequest(entry *logrus.Entry, status int) {
	switch {
	case status >= 500:
		entry.Error("request failed")
	case status >= 400:
		entry.Warn("client error")
	default:
		entry.Info("request handled")
	}
}
