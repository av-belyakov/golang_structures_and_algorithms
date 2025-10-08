package binarysearchalgorithm

import "slices"

// BinarySearch ыполняет стандартный двоичный поиск, чтобы найти целевой объект в отсортированном массиве.
// Возвращает индекс целевого объекта, если он найден, или -1, если не найден.
func BinarySearch(arr []int, target int) int {
	slices.Sort(arr)
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return -1
}

// BinarySearchRecursive выполняет двоичный поиск с использованием рекурсии.
// Возвращает индекс целевого объекта, если он найден, или -1, если не найден.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	count := len(arr)
	if count == 0 {
		return -1
	}

	if count == 1 {
		if arr[0] == target {
			return 0
		} else {
			return -1
		}
	}

	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return BinarySearchRecursive(arr, target, left, mid-1)
	}

	if arr[mid] < target {

		return BinarySearchRecursive(arr, target, mid+1, right)
	}

	return -1
}

// FindInsertPosition возвращает индекс, в который должен быть вставлен целевой объект
// для сохранения порядка сортировки массива.
func FindInsertPosition(arr []int, target int) int {
	left, right := 0, len(arr)

	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
