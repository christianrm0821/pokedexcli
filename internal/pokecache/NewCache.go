package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) CacheStruct {
	cacheMap := make(map[string]cacheEntry)

	myCache := CacheStruct{
		Cache: cacheMap,
		Mu:    &sync.Mutex{},
		Time:  interval,
	}
	myCache.reapLoop()
	return myCache
}

func (c CacheStruct) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	entry := cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Cache[key] = entry
}

func (c CacheStruct) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	val, ok := c.Cache[key]
	if !ok {
		return nil, false
	}
	return val.Val, true
}

func (c CacheStruct) reapLoop() {
	timer := time.NewTicker(c.Time)
	tmp := timer.C
	if tmp == nil {
		return
	}
	for k, v := range c.Cache {
		tmp_time := time.Since(v.CreatedAt)
		if tmp_time >= c.Time {
			delete(c.Cache, k)
		}
	}
}
