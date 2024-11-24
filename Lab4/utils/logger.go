package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

var Logger *zap.Logger

func InitLogger() {
	logWriter := &lumberjack.Logger{
		Filename:   "./logs/app.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		zapcore.AddSync(logWriter),
		zapcore.InfoLevel,
	)

	Logger = zap.New(core)

	if Logger == nil {
		log.Fatalf("Failed to initialize logger")
	}
}

func LogInfo(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

func LogError(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}
