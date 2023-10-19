package loggertests

import "testing"

import "github.com/sirgallo/logger"


var cLog *logger.CustomLog 

func init() {
	cLog = logger.NewCustomLog("TestLogger")
}

func TestLogger(t *testing.T) {
	t.Run("test debug", func(t *testing.T) {
		cLog.Debug("this is a debug message.")
	})

	t.Run("test error", func(t *testing.T) {
		cLog.Error("this is an error message.")
	})

	t.Run("test info", func(t *testing.T) {
		cLog.Info("this is an info message.")
	})

	t.Run("test warn", func(t *testing.T) {
		cLog.Warn("this is a warning message.")
	})
}
