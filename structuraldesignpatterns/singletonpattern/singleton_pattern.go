// - шаблон Одиночка (Singleton)
package structuraldesignpatterns

import "sync"

var (
	once     sync.Once
	memCache *SingletonMemoryCache = nil
)

type SingletonMemoryCache struct {
	Alloc                  uint64
	HeapSys                uint64
	HeapAlloc              uint64
	TotalAlloc             uint64
	HeapObjects            uint64
	NumberLiveObjects      uint64
	CountMemoryReturned    uint64
	GarbagecollectorMemory uint64
}

func NewSingletonMemoryCache() *SingletonMemoryCache {
	once.Do(func() {
		memCache = new(SingletonMemoryCache)
	})

	return memCache
}
