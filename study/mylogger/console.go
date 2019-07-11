package mylogger

import (
	"fmt"
	"os"
	"time"
)

// 往终端打印日志

// 定义ConsoleLogger终端日志结构体
type ConsoleLogger struct {
	level Level
}

// 终端日志结构体的构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	cLog := &ConsoleLogger{
		level: logLevel,
	}
	return cLog
}

// 将公用的记录日志的功能封装成一个单独的方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}
	// 用户传进来的日志字符串
	msg := fmt.Sprintf(format, args...)

	// 日志格式：[时间][日志级别][文件：行号][函数名] 日志信息
	nowStr := time.Now().Format("2006-01-02 05:20:00.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s][%s：%d] [%s] %s", nowStr, logLevelStr, fileName, line, funcName, msg)

	// 利用fmt包将日志打印
	_, _ = fmt.Fprintln(os.Stdin, logMsg)

	////如果是error或者fatal级别的日志还应该记录到errfile
	//if level >= Error {
	//	_, _ = fmt.Fprintln(os.Stderr, logMsg)
	//}
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(Debug, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(Info, format, args...)
}

func (c *ConsoleLogger) Warning(format string, args ...interface{}) {
	c.log(Warning, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(Error, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(Fatal, format, args...)
}
