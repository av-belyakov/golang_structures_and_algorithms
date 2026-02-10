package twopointeralgorithm

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// алгоритм двойного указателя (часто используется когда надо в срезе чисел
// найти два числа которые в сумме дают какое то определённое третье число)
// подробнее https://algodaily.com/lessons/using-the-two-pointer-technique/go

func twoSum(arr []int, target int) []int {
	sort.Ints(arr)
	left := 0
	right := len(arr) - 1

	for left < right {
		if arr[left]+arr[right] == target {
			return []int{arr[left], arr[right]}
		} else if arr[left]+arr[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {
	arr := []int{1, 3, 2, 9, 4, 8, 9}
	target := 11
	result := twoSum(arr, target)

	assert.Equal(t, result[0], 2)
	assert.Equal(t, result[0], 9)

	fmt.Println("Result:", result)
}
