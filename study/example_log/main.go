package main

import (
	"golang/study/mylogger"
)

func main() {
	log := mylogger.NewFileLogger("./", "test.log")
	log.Debug("日志库实现1")
}
