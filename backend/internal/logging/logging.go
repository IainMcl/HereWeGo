package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/IainMcl/HereWeGo/internal/file"
	"github.com/IainMcl/HereWeGo/internal/settings"
)

type Level int

var (
	F        *os.File
	LogLevel = DEBUG

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  string
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	LogLevel = getLogLevel()
	filePath := getLogFilePath()
	fileName := getLogFileName()
	if settings.AppSettings.LogToConsole {
		F = os.Stdout
	} else {
		F, err = file.MustOpen(fileName, filePath)
		if err != nil {
			log.Fatalf("logging.Setup err: %v", err)
		}
	}
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func getLogLevel() Level {
	switch strings.ToLower(settings.AppSettings.LogLevel) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return DEBUG
	}
}

func Debug(v ...interface{}) {
	if LogLevel > DEBUG {
		return
	}
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	if LogLevel > INFO {
		return
	}
	setPrefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	if LogLevel > WARNING {
		return
	}
	setPrefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	if LogLevel > ERROR {
		return
	}
	setPrefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	if LogLevel > FATAL {
		return
	}
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
