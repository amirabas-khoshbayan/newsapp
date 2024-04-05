package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var ZapLogger *zap.Logger
var once = sync.Once{}

func init() {
	once.Do(func() {
		ZapLogger, _ = zap.NewProduction()
		defaultLogLevel := zapcore.InfoLevel

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:  "./logs/log.json",
			LocalTime: false,
			MaxSize:   10, // megabytes
			//MaxBackups: 10,
			MaxAge: 30, // days
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			// write to stdout as well as log files
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer), zap.NewAtomicLevelAt(defaultLogLevel))

		ZapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	})
}
