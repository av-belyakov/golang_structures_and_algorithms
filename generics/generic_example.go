package generics

// функция для приска любого сравнимого значения в в срезе
func SearchInSlice[T comparable](elem T, l []T) (T, bool) {
	var result T

	for _, v := range l {
		if elem == v {
			result = v

			return result, true
		}
	}

	return result, false
}

// функция выполняющая сбор всех ключей в срез
func GetKeysFromMap[T comparable, V any](list map[T]V) []T {
	result := make([]T, 0, len(list))

	for k := range list {
		result = append(result, k)
	}

	return result
}

type MyNumberType interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64
}

func MySummation[T MyNumberType](numbers []T) T {
	var result T

	for _, v := range numbers {
		result += v
	}

	return result
}
