package mylogger

import "strings"

// 我的日志库文件

// Level 是一个自定义类型 代表日志级别
type Level uint16

// 定义具体的日志级别常量
const (
	Debug Level = iota
	Info
	Warning
	Error
	Fatal
)

// 根据传进来的日志级别，获取对应的字符串
func getLevelStr(level Level) string {
	switch level {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// 根据用户传入的字符串类型日志级别，解析出对应的Level
func parseLogLevel(levelStr string) Level {
	// 字符串转小写
	levelStr = strings.ToLower(levelStr)
	switch levelStr {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warning":
		return Warning
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Debug
	}
}
