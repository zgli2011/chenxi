package middleware

import (
	"bytes"
	"chenxi/initialize"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Entry

func LoggerToFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get(initialize.TraceIDKey) == "" {
			ctx.Request.Header.Set(initialize.TraceIDKey, uuid.NewV4().String())
		}
		Logger = initialize.Logger.WithFields(logrus.Fields{
			"appName":             "transaction_orchestration",
			"host":                initialize.Config.System.LocalHostName,
			"ip":                  initialize.Config.System.LocalIP,
			initialize.TraceIDKey: ctx.Request.Header.Get(initialize.TraceIDKey),
		})
		//日志格式
		Logger = Logger.WithFields(map[string]interface{}{
			initialize.TraceIDKey: ctx.Request.Header.Get(initialize.TraceIDKey),
		})
		Logger.WithContext(ctx).Infof("--> %s | %s",
			ctx.Request.Method,
			ctx.Request.RequestURI,
		)
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		if len(body) > 0 {
			Logger.Infof("%s", body)
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		startTime := time.Now()               // 开始时间
		ctx.Next()                            // 处理请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		statusCode := ctx.Writer.Status()     // 状态码
		// 日志格式
		Logger.Infof("<-- %s | %s | %s | %3d | %13v",
			ctx.Request.Method,
			ctx.Request.RequestURI,
			body,
			statusCode,
			latencyTime,
		)
	}
}
