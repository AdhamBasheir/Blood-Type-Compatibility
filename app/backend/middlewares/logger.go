package middlewares

import (
	"blood-type-compatibility/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		latency := helpers.MeasureLatency(ctx.Next)
		fields := getBaseRequestLogFields(ctx)
		fields["latency"] = latency.String()

		entry := logrus.WithFields(fields)
		logRequest(entry, ctx.Writer.Status())
	}
}
