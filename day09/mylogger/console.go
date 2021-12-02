package mylogger

import (
	"fmt"
	"time"
)

type consoleLogger struct {
	LEVEL logLevel
}

func NewLog(levelStr string) consoleLogger {
	level := transformStrToLevel(levelStr)
	return consoleLogger{
		LEVEL: level,
	}
}

func (l consoleLogger)enable(level logLevel)bool{
	return l.LEVEL <= level
}

func (l consoleLogger)Debug(s string){
	if l.enable(DEBUG) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"),funcName,fileName,lineNo, s)
	}
}

func (l consoleLogger)Info(s string){
	if l.enable(INFO) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"),funcName,fileName,lineNo, s)
	}
}

func (l consoleLogger)Warning(s string){
	if l.enable(WARNING) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"),funcName,fileName,lineNo, s)
	}
}
func (l consoleLogger)Error(s string) {
	if l.enable(ERROR){
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"),funcName,fileName,lineNo, s)
	}
}
func (l consoleLogger)Fatal(s string){
	if l.enable(FATAL) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Printf("[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"),funcName,fileName,lineNo, s)
	}
}


