package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 日志写入文件

// FileLogger 文件日志结构体
type FileLogger struct {
	level    Level
	fileName string
	filePath string
	file     *os.File
	errFile  *os.File
}

// 文件日志结构体的构造函数
func NewFileLogger(levelStr, filePath, fileName string) *FileLogger {
	logLevel := parseLogLevel(levelStr)
	fLog := &FileLogger{
		level:    logLevel,
		fileName: fileName,
		filePath: filePath,
	}
	fLog.initFile()
	return fLog
}

// 初始化文件
func (f *FileLogger) initFile() {
	logName := path.Join(f.filePath, f.fileName)

	// 生成日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败, %v", logName, err))
	}
	f.file = fileObj

	// 生成错误日志文件
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件%s失败, %v", errLogName, err))
	}

	f.errFile = errFileObj

}

// 将公用的记录日志的功能封装成一个单独的方法
func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	// 用户传进来的日志字符串
	msg := fmt.Sprintf(format, args...)

	// 日志格式：[时间][日志级别][文件：行号][函数名] 日志信息
	nowStr := time.Now().Format("2006-01-02 05:20:00.000")
	fileName, line, funcName := getCallerInfo(3)
	logLevelStr := getLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s][%s：%d] [%s] %s", nowStr, logLevelStr, fileName, line, funcName, msg)

	// 利用fmt包将日志写入文件
	_, _ = fmt.Fprintln(f.file, logMsg)

	//如果是error或者fatal级别的日志还应该记录到errfile
	if level >= Error {
		_, _ = fmt.Fprintln(f.errFile, logMsg)
	}
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(Debug, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(Info, format, args...)
}

func (f *FileLogger) Warning(format string, args ...interface{}) {
	f.log(Warning, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(Error, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(Fatal, format, args...)
}
