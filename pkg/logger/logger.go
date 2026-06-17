package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

func InitLogger() error {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.LevelKey = "level"
	config.EncoderConfig.MessageKey = "message"

	var err error
	Log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}