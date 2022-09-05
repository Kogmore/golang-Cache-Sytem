package main

import (
	"golang-Cache-Sytem/cache_server"
	"log"
)

func main() {
	newCache := cache_server.NewCache() //变量改个名字 是因为与导入的包名冲突

	//设置缓存内存大小
	ok := newCache.SetMaxMemory("5B")
	if ok {
		log.Println("Set Max Memory 设置成功")
	} else {
		log.Println("Set Max Memory 设置失败")
	}

	//设置缓存项 key：int val：1	无过期时间
	ok = newCache.Set("int", 1)
	if ok {
		log.Printf("Set cache: key:%s 成功", "int")
	} else {
		log.Println("Set cache 失败")
	}

	//设置缓存项 key：bool val：false 无过期时间
	ok = newCache.Set("bool", false)
	if ok {
		log.Printf("Set cache: key:%s 成功", "bool")
	} else {
		log.Println("Set cache 失败")
	}

	//设置缓存项 key：data val：map[string]interface{}{"a": 1} 无过期时间
	ok = newCache.Set("data", map[string]interface{}{"a": 1})
	if ok {
		log.Printf("Set cache: key:%s 成功", "data")
	} else {
		log.Println("Set cache 失败")
	}

	var date interface{}
	//获取缓存key为int的值
	date, ok = newCache.Get("int")
	if ok {
		log.Printf("Get cache: val:%v 成功", date)
	} else {
		log.Println("Get cache 失败")
	}

	////删除缓存key为int的键值对
	//ok = newCache.Del("int")
	//if ok {
	//	log.Printf("Del cache: val:%v 成功", date)
	//} else {
	//	log.Println("Get cache 失败")
	//}
	//
	////清空所有缓存项
	//ok = newCache.Flush()
	//if ok {
	//	log.Println("Flush cache 成功")
	//} else {
	//	log.Println("Flush cache 失败")
	//}

	//查询当前缓存有多少键值对
	keys := newCache.Keys()
	log.Printf("get keys:%d", keys)
}
