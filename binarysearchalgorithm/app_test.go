package binarysearchalgorithm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/binarysearchalgorithm"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	t.Run("Обычный бинарный поиск", func(t *testing.T) {
		assert.Equal(t, binarysearchalgorithm.BinarySearch(arr, 4), 3)
	})

	t.Run("Рекурсивный бинарный поиск", func(t *testing.T) {
		assert.Equal(t, binarysearchalgorithm.BinarySearchRecursive(arr, 6, 0, len(arr)), 5)
	})

	t.Run("Поиск индекса для вставки", func(t *testing.T) {
		assert.Equal(t, binarysearchalgorithm.FindInsertPosition(arr, 10), 9)
	})
}
