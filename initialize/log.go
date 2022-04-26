package initialize

import (
	"chenxi/utils"
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const (
	TraceIDKey = "traceID"
)

var Logger *logrus.Logger

type LogOptions struct {
	Path         string `yaml:"path" json:"path"`
	Level        string `yaml:"level" json:"level"`
	MaxAge       int    `yaml:"max_age" json:"max_age"`
	RotationTime int    `yaml:"rotation_time" json:"rotation_time"`
}

func NewLog() *logrus.Logger {
	log_level := map[string]int{"trace": 6, "debug": 5, "info": 4, "warn": 3, "error": 2, "fatal": 1, "panic": 0}
	log_config := Config.Log
	if _, ok := log_level[log_config.Level]; !ok {
		log_config.Level = "info"
	}

	// 检查日志文件
	logFile := log_config.Path
	if ok := utils.CheckDirOrFileExist(logFile); !ok {
		if _, err := os.Create(logFile); err != nil {
			log.Panic("日志文件创建失败:" + err.Error())
		}
	}

	Logger = logrus.New() //实例化
	// Logger.SetOutput(os.Stdout)                                //设置输出到控制台
	Logger.SetReportCaller(true)                               //设置打印caller
	Logger.SetLevel(logrus.Level(log_level[log_config.Level])) //设置日志级别

	logWriter, _ := rotatelogs.New( // 设置 rotatelogs
		logFile+".%Y%m%d.log",                                                         // 分割后的文件名称
		rotatelogs.WithLinkName(logFile),                                              // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Duration(log_config.MaxAge)*time.Hour),             // 设置最大保存时间(7天)
		rotatelogs.WithRotationTime(time.Duration(log_config.RotationTime)*time.Hour), // 设置日志切割时间间隔(1天)
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05.000Z",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyFunc:  "logger",
			logrus.FieldKeyFile:  "file",
		},
	})
	Logger.AddHook(lfHook)
	return Logger
}

// 非web程序使用的日志打印,使用了上下文，保证一个业务都能有唯一的trace_id
type (
	traceIDKey struct{}
)

func TransContext(c *gin.Context) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, traceIDKey{}, c.Request.Header.Get(TraceIDKey))
	return ctx
}

func NewTraceIDContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, traceIDKey{}, uuid.NewV4().String())
	return ctx
}

func WithContext(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		ctx = NewTraceIDContext()
	}

	fields := map[string]interface{}{
		"appName":  "transaction_orchestration",
		"host":     Config.System.LocalHostName,
		"ip":       Config.System.LocalIP,
		TraceIDKey: ctx.Value(traceIDKey{}),
	}

	return Logger.WithContext(ctx).WithFields(fields)
}
