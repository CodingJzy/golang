package mylogger

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
