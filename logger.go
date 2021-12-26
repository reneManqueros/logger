package logger

import (
	"log"
	"os"
)

const FATAL = "FATAL"
const ERROR = "ERROR"
const WARN = "WARN"
const INFO = "INFO"
const DEBUG = "DEBUG"

var levelPriority = map[string]int{
	"DEBUG": 0,
	"INFO":  1,
	"WARN":  2,
	"ERROR": 3,
	"FATAL": 4,
}

type Logger struct {
	Level string
}

var LogLevel = ""

func (l *Logger) Fatal(v ...interface{}) {
	l.logMessages(FATAL, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.logMessages(ERROR, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.logMessages(WARN, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.logMessages(DEBUG, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.logMessages(INFO, v...)
}

func (l *Logger) logMessages(level string, v ...interface{}) {
	if LogLevel != "" {
		l.Level = LogLevel
	}

	if l.Level == "" {
		l.Level = INFO
	}
	loggerPriority := levelPriority[l.Level]
	messagePriority := levelPriority[level]

	if level == l.Level || messagePriority >= loggerPriority {
		log.SetOutput(os.Stdout)
		messages := []interface{}{
			level,
			v,
		}
		log.Println(messages...)
	}
}
