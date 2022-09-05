package util

import (
	"testing"
)

func TestParseSize(t *testing.T) {
	size := "100MB"
	if num, numStr := ParseSize(size); num != 104857600 && numStr != "100MB" {
		t.Errorf("set cache size err")
	}
}

func TestGetValSize(t *testing.T) {
	val := "张三李四王五"
	if size := GetValSize(val); size != 20 {
		t.Errorf("get value size err:%d", size)
	}
}
