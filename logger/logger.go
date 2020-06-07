package logger

import (
	"go.uber.org/zap"
)

// Log es el logger a usarse en la aplicación
var Log *zap.SugaredLogger

// NewLogger devuelve un nuevo logger
func NewLogger() *zap.SugaredLogger {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	Log = l.Sugar()
	return Log
}
