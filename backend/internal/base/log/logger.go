package log

import (
	"go.uber.org/zap"
)

var zapLogger *zap.Logger

func init() {
	var err error
	zapLogger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

// Error logs an error message with fields.
func Error(msg string, fields ...zap.Field) {
	zapLogger.Error(msg, fields...)
}

// Errorf logs an error message using a format specifier.
func Errorf(format string, args ...interface{}) {
	zapLogger.Sugar().Errorf(format, args...)
}

// Info logs an informational message with fields.
func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}

// Debug logs a debug message with fields.
func Debug(msg string, fields ...zap.Field) {
	zapLogger.Debug(msg, fields...)
}

// Warn logs a warning message with fields.
func Warn(msg string, fields ...zap.Field) {
	zapLogger.Warn(msg, fields...)
}
