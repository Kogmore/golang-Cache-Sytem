package util

import (
	"encoding/json"
	"golang-Cache-Sytem/cache_const"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ParseSize(size string) (int64, string) {
	var byteNum int64 = 0
	//解析并返回一个正则表达式
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Println("ParseSize() err:", err)
		return 0, ""
	}
	//获取单位
	unit := re.ReplaceAllString(size, "")
	//获取大小
	num, err := strconv.ParseInt(strings.Replace(size, unit, "", 1), 10, 64)
	if err != nil {
		log.Println("ParseSize() err:", err)
		num = 500
		byteNum = num * cache_const.MB

		return byteNum, "500MB"
	}
	//判断用户输入内存大小是否为0
	if num == 0 { //为0  则默认大小为500MB
		//日志
		log.Println("ParseSize() 内存大小不可设置为0")
		num = 500
		byteNum = num * cache_const.MB

		return byteNum, "500MB"
	}
	//把单位字母全部改成大写
	unit = strings.ToUpper(unit)
	//判断单位是那种 并计算出最终要设置的内存大小
	switch unit {
	case "B":
		byteNum = num
	case "KB":
		byteNum = num * cache_const.KB
	case "MB":
		byteNum = num * cache_const.MB
	case "GB":
		byteNum = num * cache_const.GB
	case "TB":
		byteNum = num * cache_const.TB
	default: //单位输入错误  则默认大小为500MB
		log.Println("ParseSize 单位仅支持 B、KB、MB、GB、TB")
		num = 500
		byteNum = num * cache_const.MB
		return byteNum, "500MB"
	}

	return byteNum, size
}

func GetValSize(val interface{}) int64 {
	//获取val 内存占用的字节数
	byte, err := json.Marshal(val)
	if err != nil {
		log.Println("GetValSize json.Marshal err:", err)
	}
	valSize := len(byte)
	return int64(valSize)
}
