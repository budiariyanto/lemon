package logger

import (
	"io"
	"lemon/exceptions"
	"log"
	"os"
)

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type Level int

func (level Level) String() string {
	switch level {
	case DebugLevel:
		return "[DEBUG]"
	case InfoLevel:
		return "[INFO]"
	case WarnLevel:
		return "[WARNING]"
	case ErrorLevel:
		return "[ERROR]"
	case FatalLevel:
		return "[FATAL]"
	case PanicLevel:
		return "[PANIC]"
	}

	return "[UNKNOWN]"
}

type Logger struct {
	stdLogger *log.Logger
	LogLevel  Level
}

func newLogger(output io.Writer) *Logger {
	logger := new(Logger)
	logger.stdLogger = log.New(output, DebugLevel.String(), log.LstdFlags|log.Llongfile)
	logger.LogLevel = DebugLevel

	return logger
}

func NewStdLogger() *Logger {
	return newLogger(os.Stderr)
}

func NewFileLogger(fullFilePath string) (*Logger, error) {
	file, err := os.OpenFile(fullFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

	if err != nil {
		return NewStdLogger(), exceptions.LoggerFallbackWarning
	}

	return newLogger(file), nil
}

func (logger *Logger) log(level Level, message string) {
	if logger.LogLevel >= level {
		logger.stdLogger.SetPrefix(level.String())

		if level == FatalLevel {
			logger.stdLogger.Fatalln(message)
		} else if level == PanicLevel {
			logger.stdLogger.Panicln(message)
		} else {
			logger.stdLogger.Println(message)
		}
	}
}

func (logger *Logger) Debug(message string) {
	if logger.LogLevel >= DebugLevel {
		logger.log(DebugLevel, message)
	}
}

func (logger *Logger) Info(message string) {
	if logger.LogLevel >= InfoLevel {
		logger.log(InfoLevel, message)
	}
}

func (logger *Logger) Warning(message string) {
	if logger.LogLevel >= WarnLevel {
		logger.log(WarnLevel, message)
	}
}

func (logger *Logger) Error(message string) {
	if logger.LogLevel >= ErrorLevel {
		logger.log(ErrorLevel, message)
	}
}

func (logger *Logger) Fatal(message string) {
	if logger.LogLevel >= FatalLevel {
		logger.log(FatalLevel, message)
	}
}

func (logger *Logger) Panic(message string) {
	if logger.LogLevel >= PanicLevel {
		logger.log(PanicLevel, message)
	}
}
