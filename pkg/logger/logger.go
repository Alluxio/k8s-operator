package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	zapLog      *zap.Logger
	sugarZapLog *zap.SugaredLogger
)

func init() {
	var err error
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	sugarZapLog = zapLog.Sugar()
}

func Infof(format string, args ...interface{}) {
	sugarZapLog.Infof(format, args)
}

func Errorf(format string, args ...interface{}) {
	sugarZapLog.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	sugarZapLog.Fatalf(format, args)
}
