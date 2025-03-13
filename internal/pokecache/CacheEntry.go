package pokecache

import (
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}
