package main


// 日志库需求分析
// 1.支持往不同的地方输出日志
// 2.日志要分级别
//  2.1 debug 2.2info 2.3 warning 2.4 error 2.5 fatal
// 3.日志要支持开关控制  想要让Debug级别的往下的输出那么就传入debug
// 4.日志要有 ,时间,行号,文件名,日志级别,日志信息
// 5.日志文件要切割

import (
	"github.com/day09/mylogger"
	"time"
)

func main() {
	//log := mylogger.NewLog("error")
	log := mylogger.NewFileLogger("Info","./","zhouling.log",1024 * 1024 * 10)
	for {
		log.Debug("这是一条Debug日志!")
		log.Info("这是一条Info日志!")
		log.Warning("这是一条Warning日志!")
		log.Error("这是一条Error日志!")
		log.Fatal("这是一条Fatal日志!")
		time.Sleep(time.Second * 3)
	}
	log.MyClose()
}
