package pokecache

import (
	"sync"
	"time"
)

type CacheStruct struct {
	Cache map[string]cacheEntry
	Mu    *sync.Mutex
	Time  time.Duration
}
