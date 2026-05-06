package withshards

import (
	"hash/fnv"
	"sync"
)

type InMemoryCacheWithShards struct {
	shards []Shard
}

type Shard struct {
	data map[string]string
	mtx  sync.RWMutex
}

func New(countShards int) *InMemoryCacheWithShards {
	shards := make([]Shard, 0, countShards)

	for range countShards {
		shards = append(shards, Shard{data: map[string]string{}})
	}

	return &InMemoryCacheWithShards{
		shards: shards,
	}
}

func (s *InMemoryCacheWithShards) Get(key string) (string, bool) {
	shardId := getHash(key) % len(s.shards)

	return s.shards[shardId].Get(key)
}

func (s *InMemoryCacheWithShards) Set(key, value string) {
	shardId := getHash(key) % len(s.shards)

	println("shardId: ", shardId, " key: ", key, " value: ", value)

	s.shards[shardId].Set(key, value)
}

func (s *Shard) Get(key string) (string, bool) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()

	v, ok := s.data[key]
	return v, ok
}

func (s *Shard) Set(key, value string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.data[key] = value
}

func getHash(key string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))

	return int(h.Sum32())
}
