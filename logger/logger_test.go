package logger_test

import (
	"testing"

	"github.com/santiagoh1997/challenge/logger"
)

func TestNewLogger(t *testing.T) {
	l := logger.NewLogger()
	if l == nil {
		t.Errorf("NewLogger got %v, want *zap.SugaredLogger", l)
	}
}
