package logger

import (
	"fmt"
	"testing"
)

func TestFileLogger(t *testing.T) {
	_, err := NewFileLogger("/data/development/gopkg/src/lemon/logger/mylog.log")

	if err != nil {
		// actually if file logger cannot be created, logger will fallbacked using
		// console to outputing the log. But for the sake of this test, it will be
		// considered as error and test will failed.
		t.Error("Logger file cannot be created.")
	}
}

func TestLeveledLogger(t *testing.T) {
	levels := []Level{DebugLevel, InfoLevel, WarnLevel, ErrorLevel}
	logger := NewStdLogger()

	for _, level := range levels {
		fmt.Println("Entering test", level, "level...")
		logger.LogLevel = level
		logger.Debug("Test debug output...")
		logger.Info("Test info output...")
		logger.Warning("Test warning output...")
		logger.Error("Test error output...")
	}

	// fatal and panic level cannot be outputed because it will cause test to fail
	// and exit
}
