package main

import (
	"golang/study/mylogger"
)

func main() {
	log := mylogger.NewFileLogger("info", "./", "test.log")
	log.Debug("日志库实现debug")
	log.Info("日志库实现info")
	log.Error("日志库实现error")
}
