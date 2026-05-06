package withshards_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/av-belyakov/golang_structures_and_algorithms/inmemorycache/withshards"
)

func TestInMemoryCacheWithShards(t *testing.T) {
	t.Run("Test 1. In-memory cache with shards", func(t *testing.T) {
		cache := withshards.New(10)

		for range 100 {
			cache.Set(gofakeit.ID(), gofakeit.FarmAnimal())
		}
	})
}
