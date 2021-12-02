package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志

type fileLogger struct {
	LEVEL       logLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	errFileName string
	maxFileSize int64
	fileObj     *os.File
	fileErr     *os.File
}

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *fileLogger {
	loglevel := transformStrToLevel(levelStr)
	f1 :=  &fileLogger{
		LEVEL:       loglevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	// 按照文件路径和文件名将文件打开
	err := f1.initFile()
	if err != nil {
		panic(err)
	}
	return f1
}

func (f *fileLogger) initFile() error {
	fileFullName := path.Join(f.filePath,f.fileName)
	fileObj,err := os.OpenFile(
		fileFullName,
		os.O_APPEND | os.O_CREATE | os.O_WRONLY,
		0644,
	)
	if err != nil{
		fmt.Printf("open file failed err : %v",err)
		return err
	}
	errFileObj,err := os.OpenFile(fileFullName + ".err",os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)
	if err != nil {
		fmt.Printf("open err log file failed err : %v",err)
		return err
	}
	// 日志都已经打开了
	f.fileObj = fileObj
	f.fileErr = errFileObj
	return nil
}

func (f *fileLogger) enable(level logLevel) bool {
	return f.LEVEL <= level
}

func (f *fileLogger) checkSize(fileObj *os.File) bool {
	fileInfo ,err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get fileInfo failed err = %v\n",err)
		return false
	}
	size := fileInfo.Size()
	// 如果当前文件大小大于等于 日志文件的最大值 就应该切割 返回true
	if size >= f.maxFileSize {
		return true
	}
	return false
}

//func (f *fileLogger) splitFile(file *os.File)(*os.File ,error){}

func (f *fileLogger) Debug(s string) {
	if f.enable(DEBUG) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		if f.checkSize(f.fileObj){
			// 需要切割日志文件
			// 1.关闭当前的日志文件
			f.fileObj.Close()
			// 2.rename 备份一下 xx.log -> xx.log.bak202106251019
			nowStr := time.Now().Format("20060102 15:04:05")
			logName := path.Join(f.filePath,f.fileName)
			var newLogName string
			newLogName = path.Join(logName,nowStr)
			os.Rename(logName,newLogName)
			// 3.打开一个新的文件
			fileObj,err := os.OpenFile(logName,os.O_APPEND | os.O_CREATE | os.O_WRONLY,0644)
			if err != nil{
				fmt.Printf("open new log file failed err : %v\n",err)
			}
			// 4.将打开的新日志文件对象赋值给f.fileObj
			f.fileObj = fileObj
		}
		fmt.Fprintf(f.fileObj,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
	}
}

func (f *fileLogger) Info(s string) {
	if f.enable(INFO) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Fprintf(f.fileObj,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
	}
}

func (f *fileLogger) Warning(s string) {
	if f.enable(WARNING) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Fprintf(f.fileObj,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
	}
}
func (f *fileLogger) Error(s string) {
	if f.enable(ERROR) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Fprintf(f.fileObj,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
		fmt.Fprintf(f.fileErr,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
	}
}
func (f *fileLogger) Fatal(s string) {
	if f.enable(FATAL) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(2)
		fmt.Fprintf(f.fileObj,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
		fmt.Fprintf(f.fileErr,"[%s] [%s : %s : %d]  %s\n", now.Format("2006-01strconv_demo-02 15:04:05"), funcName, fileName, lineNo, s)
	}
}

func (f *fileLogger) MyClose(){
	f.fileObj.Close()
	f.fileErr.Close()
}