package cache

import (
	"log"
	"testing"
	"time"
)

func TestMyCache_SetMaxMemory(t *testing.T) {
	cache := NewCache()
	if ok := cache.SetMaxMemory("-100MB"); !ok {
		t.Errorf("ser max memory fail,%v", ok)
	} else {
		log.Println(ok)
	}
}

func TestMyCache_Set(t *testing.T) {
	cache := NewCache()
	if ok := cache.Set("int", 1, time.Second); !ok {
		t.Errorf("set cache key:int val:1 fail,%v", ok)
	} else {
		log.Println(ok)
	}
}

func TestMyCache_Get(t *testing.T) {
	cache := NewCache()
	if _, ok := cache.Get("int"); ok {
		t.Errorf("get cache key:int,%v", ok)
	} else {
		log.Println(ok)
	}
}

func TestMyCache_Del(t *testing.T) {
	cache := NewCache()
	if ok := cache.Del("int"); ok {
		t.Errorf("del cache key:int,%v", ok)
	} else {
		log.Println(ok)
	}
}

func TestMyCache_Exists(t *testing.T) {
	cache := NewCache()
	if ok := cache.Exists("int"); ok {
		t.Errorf("exists cache key:int,%v", ok)
	} else {
		log.Println(ok)
	}
}

func TestMyCache_Flush(t *testing.T) {
	cache := NewCache()
	if ok := cache.Flush(); !ok {
		t.Errorf("flush cache,%v", ok)
	} else {
		log.Println(ok)
	}
}

func TestMyCache_Keys(t *testing.T) {
	cache := NewCache()
	if keyNum := cache.Keys(); keyNum != 0 {
		t.Errorf("flush cache,%d", keyNum)
	} else {
		log.Println(keyNum)
	}
}
