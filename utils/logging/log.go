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

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
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
	filePath := getLogFilePath()
	fileName := getLogFileName()
	F, err = openLogFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{})  {
	setPrefix(DEBUG)
	logger.Print(v)
}

func Info(v ...interface{})  {
	setPrefix(INFO)
	logger.Print(v)
}

func Warning(v ...interface{})  {
	setPrefix(WARNING)
	logger.Print(v)
}

func Error(v ...interface{})  {
	setPrefix(ERROR)
	logger.Print(v)
}

func Fatal(v ...interface{})  {
	setPrefix(FATAL)
	logger.Print(v)
}

func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s]:[%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}