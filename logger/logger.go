package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

const (
	defaultFilePath        = "./logs/logs.json"
	defaultUseLocalTime    = false
	defaultFileMaxSizeInMB = 10
	defaultFileAgeInDays   = 30
)

type Config struct {
	FilePath         string `yaml:"file_path"`
	UseLocalTime     bool   `yaml:"use_local_time"`
	FileMaxSizeInMB  int    `yaml:"file_max_size_in_mb"`
	FileMaxAgeInDays int    `yaml:"file_max_age_in_days"`
}

var ZapLogger *zap.Logger
var once = sync.Once{}

func Init(cfg Config) {
	once.Do(func() {
		ZapLogger, _ = zap.NewProduction()
		defaultLogLevel := zapcore.InfoLevel

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:  cfg.FilePath,
			LocalTime: cfg.UseLocalTime,
			MaxSize:   cfg.FileMaxSizeInMB, // megabytes
			//MaxBackups: 10,
			MaxAge: cfg.FileMaxAgeInDays, // days
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			// write to stdout as well as log files
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), writer), zap.NewAtomicLevelAt(defaultLogLevel))

		ZapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	})
}
