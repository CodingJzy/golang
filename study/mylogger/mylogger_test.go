package mylogger

import "testing"

// 单元测试

func TestConst(t *testing.T) {
	t.Logf("类型：%T 值：%v", Debug, Debug)
	t.Logf("类型：%T 值：%v", Error, Error)
}
