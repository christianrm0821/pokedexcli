package pokecache

import (
	"sync"
	"time"
)

func NewCache(timeout time.Duration) CacheStruct {
	cacheMap := make(map[string]cacheEntry)

	myCache := CacheStruct{
		Cache: cacheMap,
		Mu:    &sync.Mutex{},
		Time:  timeout,
	}
	return myCache
}

func (c CacheStruct) Add(key string, val []byte) {
	entry := cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Cache[key] = entry
}

func (c CacheStruct) Get(key string) ([]byte, bool) {
	val, ok := c.Cache[key]
	if !ok {
		return nil, false
	}
	return val.Val, true
}
