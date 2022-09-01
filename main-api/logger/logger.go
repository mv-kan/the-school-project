package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() Logger {
	zapConfig := zap.NewProductionEncoderConfig()
	zapConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(zapConfig)

	f, err := os.OpenFile("./main-api.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileEncoder := zapcore.NewJSONEncoder(zapConfig)

	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(f), defaultLogLevel),
	)
	zapLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.InfoLevel))

	return &logger{zapLogger: zapLogger}
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Error(msg string)
	Fatal(msg string)
}

type logger struct {
	zapLogger *zap.Logger
}

func (l *logger) Debug(msg string) {
	l.zapLogger.Debug(msg)
}

func (l *logger) Info(msg string) {
	l.zapLogger.Info(msg)
}

func (l *logger) Error(msg string) {
	l.zapLogger.Error(msg)
}

func (l *logger) Fatal(msg string) {
	l.zapLogger.Fatal(msg)
}
