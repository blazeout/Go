package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type logLevel uint16

const (
	// UNKNOWN debug 定义日志级别

	UNKNOWN logLevel = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
)

// consoleLogger 日志对象

// NewLog 日志构造函数返回一个logger对象

func transformStrToLevel(level string) logLevel {
	s := strings.ToLower(level)
	switch s {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warning":
		return WARNING
	case "error":
		return ERROR
	case "fatal":
		return FATAL
	default:
		return UNKNOWN
	}
}

func getInfo(n int) (funcName string, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(n)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fileName = path.Base(file)
	return funcName, fileName, lineNo
}
