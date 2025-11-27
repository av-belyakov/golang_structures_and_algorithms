package slicemanipulation

import (
	"reflect"
	"slices"
)

// SliceContainsElement проверяет наличие элемента в срезе
func SliceContainsElement[T comparable](elem T, list []T) (int, bool) {
	for k, v := range list {
		if v == elem {
			return k, true
		}
	}

	return -1, false
}

// SliceContainsElementFunc проверяет наличие элемента в срезе
func SliceContainsElementFunc[L any](list []L, f func(n int) bool) (int, bool) {
	for k := range list {
		if f(k) {
			return k, true
		}
	}

	return -1, false
}

// ReplacingSlice заменяет элементы в срезе
func ReplacingSlice[T comparable](current, new reflect.Value) (list reflect.Value, ok bool) {
	if reflect.DeepEqual(current, new) {
		return list, false
	}

	currentTags, okCurr := current.Interface().([]T)
	newTags, okNew := new.Interface().([]T)
	if !okCurr || !okNew {
		return list, false
	}

	if !current.CanSet() {
		return list, false
	}

	list = reflect.ValueOf(SliceJoinUniq(currentTags, newTags))

	return list, true
}

// SliceJoinUniq создает новый срез состоящий только из уникальных
// значений срезов sliceOne и sliceTwo
func SliceJoinUniq[T comparable](sliceOne, sliceTwo []T) []T {
	numOne, numTwo := len(sliceOne), len(sliceTwo)

	if numOne == 0 && numTwo == 0 {
		return []T{}
	}

	if numOne == 0 {
		return sliceTwo
	}

	if numTwo == 0 {
		return sliceOne
	}

	newList := make([]T, 0, numOne+numTwo)

	for _, v := range append(sliceOne, sliceTwo...) {
		if SliceElemIsExist(v, newList) {
			continue
		}

		newList = append(newList, v)
	}

	return newList
}

// SliceElemIsExist проверяет наличие элемента в срезе
func SliceElemIsExist[T comparable](item T, list []T) bool {
	return slices.Contains(list, item)
}

// RemoveElement удаляет элемент из среза
func RemoveElement[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}
