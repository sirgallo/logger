package loggertests

import "testing"

import "github.com/sirgallo/logger"


func TestLogger(t *testing.T) {

	t.Run("test debug", func(t *testing.T) {
		dLog := logger.NewCustomLog("TestDebugLogger")
		dLog.Debug("this is a debug message.")
	})

	t.Run("test error", func(t *testing.T) {
		eLog := logger.NewCustomLog("TestErrorLogger")
		eLog.Error("this is an error message.")
	})

	t.Run("test info", func(t *testing.T) {
		iLog := logger.NewCustomLog("TestInfoLogger")
		iLog.Info("this is an info message.")
	})

	t.Run("test warn", func(t *testing.T) {
		wLog := logger.NewCustomLog("TestWarnLogger")
		wLog.Warn("this is a warning message.")
	})
}
