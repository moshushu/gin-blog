package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := getLogFilePath()
	fileName := getLogFileName()

	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	// LstdFlags 日志格式属性
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}
func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v...)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v...)
}

// 自定义日志前缀
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
