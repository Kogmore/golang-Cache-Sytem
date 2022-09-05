package cache_server

import (
	"golang-Cache-Sytem/cache"
	"time"
)

/*
	适配层
*/
type cacheServer struct {
	myCache cache.Cache
}

func NewCache() *cacheServer {
	return &cacheServer{
		myCache: cache.NewCache(),
	}
}

//SetMaxMemory size 是⼀个字符串。⽀持以下参数: 1KB，100KB，1MB，2MB，1GB 等
func (cs *cacheServer) SetMaxMemory(size string) bool {
	return cs.myCache.SetMaxMemory(size)
}

//Set 设置⼀个缓存项，并且在expire时间之后过期
func (cs *cacheServer) Set(key string, val interface{}, expire ...time.Duration) bool {
	expires := time.Second * 0
	if len(expire) > 0 {
		expires = expire[0]
	}
	return cs.myCache.Set(key, val, expires)
}

//Get 获取⼀个值
func (cs *cacheServer) Get(key string) (interface{}, bool) {
	return cs.myCache.Get(key)
}

//Del 删除⼀个值
func (cs *cacheServer) Del(key string) bool {
	return cs.myCache.Del(key)
}

//Exists 检测⼀个值 是否存在
func (cs *cacheServer) Exists(key string) bool {
	return cs.myCache.Exists(key)
}

//Flush 情况所有值
func (cs *cacheServer) Flush() bool {
	return cs.myCache.Flush()
}

//Keys 返回所有的key 多少
func (cs *cacheServer) Keys() int64 {
	return cs.myCache.Keys()
}
