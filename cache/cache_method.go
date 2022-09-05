package cache

import (
	"golang-Cache-Sytem/util"
	"log"
	"sync"
	"time"
)

type MyCache struct {
	maxMemory       int64                 //最大内存
	maxMemoryString string                //最大内存字符串表示  方便查询设置的最大内存为多少
	currMemory      int64                 //目前已使用内存
	values          map[string]*valExpire //用来存储缓存数据的map
	mutex           sync.RWMutex          //map非线程安全 加读写锁
}

//values map 存储的值
type valExpire struct {
	val        interface{}   //value值
	expireTime time.Time     //过期时间
	expire     time.Duration //有效时长
	size       int64         //value 大小 为了方便计算已使用内存
}

func NewCache() Cache {
	mc := &MyCache{}
	//初始化map
	mc.values = make(map[string]*valExpire)
	//定期清除过期缓存的goroutine
	go mc.clearExpiredItem()
	return mc
}

//SetMaxMemory size 是⼀个字符串。⽀持以下参数: 1KB，100KB，1MB，2MB，1GB 等
func (mc *MyCache) SetMaxMemory(size string) bool {
	//设置最大内存
	mc.maxMemory, mc.maxMemoryString = util.ParseSize(size)
	if mc.maxMemory == 0 {
		log.Println("util.ParseSize()")
		return false
	}
	return true
}

//Set 设置⼀个缓存项，并且在expire时间之后过期
func (mc *MyCache) Set(key string, val interface{}, expire time.Duration) bool {
	//声明一个valExpire的对象
	v := &valExpire{
		val:        val,                    //值
		expireTime: time.Now().Add(expire), //过期时间
		expire:     expire,                 //有效时长
		size:       util.GetValSize(val),   //内存大小
	}
	//查询是否存在 存在则删除
	mc.Del(key)
	//添加数据
	mc.mutex.Lock()
	mc.values[key] = v
	mc.currMemory += v.size
	mc.mutex.Unlock()
	//判断是否超出最大内存
	if mc.currMemory >= mc.maxMemory {
		//先删除当前添加的这个值
		mc.Del(key)
		//遍历map
		for k, value := range mc.values {
			//expire不为0 则代表设置了过期时间  expireTime过期时间在当前时间之前 表示已过期
			if value.expire != 0 && value.expireTime.Before(time.Now()) {
				mc.Del(k)
			} else {
				mc.Del(k)
			}
		}
	}
	return true
}

//Get 获取⼀个值
func (mc *MyCache) Get(key string) (interface{}, bool) {
	//查询数据
	mcVal, ok := mc.get(key)
	if ok {
		//判断缓存是否过期
		if mcVal.expire != 0 && mcVal.expireTime.Before(time.Now()) { //首先判断是否有设置过期时长  	过期时间 早于当前时间 则过期
			//过期则删除
			mc.Del(key)
			return nil, false
		} else {
			//未过期  正常返回数据
			return mcVal.val, ok
		}
	}
	//查询不到数据
	return nil, false
}

//Del 删除⼀个值
func (mc *MyCache) Del(key string) bool {
	//查询数据
	temp, ok := mc.get(key)
	if ok { //存在
		//先减去内存
		mc.mutex.Lock()
		mc.currMemory -= temp.size
		//删除数据
		delete(mc.values, key)
		mc.mutex.Unlock()
		return true
	} else {
		return false
	}
}

//Exists 检测⼀个值 是否存在
func (mc *MyCache) Exists(key string) bool {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	_, ok := mc.get(key)

	return ok
}

//Flush 清空所有值
func (mc *MyCache) Flush() bool {
	//第一种方法
	mc.values = make(map[string]*valExpire)
	////第二种方法
	//for key := range mc.values {
	//	/*
	//		使用for  range  遍历一个map
	//		然后delete 删除元素
	//		编译器会优化成mapClear函数去删除所有元素
	//		性能方面
	//		10w以内推荐第一种
	//		超过10w推荐第二种
	//	*/
	//	delete(mc.values, key)
	//}
	return true
}

//Keys 返回所有的key 多少
func (mc *MyCache) Keys() int64 {
	mc.mutex.RLock()
	defer mc.mutex.RUnlock()

	mLen := len(mc.values)
	return int64(mLen)
}

//根据key查询val
func (mc *MyCache) get(key string) (*valExpire, bool) {
	mc.mutex.RLock()
	val, ok := mc.values[key]
	mc.mutex.RUnlock()
	return val, ok
}

//用来定期删除过期key的goroutine
func (mc *MyCache) clearExpiredItem() {
	//定时器  10秒清除一次
	timeTick := time.NewTicker(time.Second * 10)
	//死循环
	for {
		select {
		case <-timeTick.C: //channel C 取到值则触发定时器
			//遍历map所有的值
			for key, val := range mc.values {
				//expire不为0 则代表设置了过期时间  expireTime过期时间在当前时间之前 表示已过期
				if val.expire != 0 && val.expireTime.Before(time.Now()) {
					mc.Del(key)
				}
			}
		}

	}
}
