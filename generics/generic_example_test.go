package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInSlice(t *testing.T) {
	//*************************
	//поиск строковых значений
	listStringOne := []string{"red", "green", "yellow", "white", "blue", "grey", "black", "orange"}

	searchStringOne := "white"
	resultStr, okStr := SearchInSlice[string](searchStringOne, listStringOne)
	assert.True(t, okStr)
	assert.Equal(t, searchStringOne, resultStr)

	searchStringTwo := "pink"
	resultStr, okStr = SearchInSlice[string](searchStringTwo, listStringOne)
	assert.False(t, okStr)
	assert.NotEqual(t, searchStringTwo, resultStr)

	//*************************
	//поиск числовых значений
	listIntOne := []int{1, 2, 6, 4, 7, 9, 11, 24, 89}
	searchIntOne := 4
	resultInt, okInt := SearchInSlice[int](searchIntOne, listIntOne)
	assert.True(t, okInt)
	assert.Equal(t, searchIntOne, resultInt)

	searchIntTwo := 99
	resultInt, okInt = SearchInSlice[int](searchIntTwo, listIntOne)
	assert.False(t, okInt)
	assert.NotEqual(t, searchIntTwo, resultInt)

	//************************
	//сбор всех ключей типа string в слайс
	mapStrAnimals := map[string]bool{"cat": true, "dog": true, "fox": false, "cow": true, "horse": true, "monkey": false}
	keysAnimals := GetKeysFromMap[string, bool](mapStrAnimals)
	assert.Equal(t, len(keysAnimals), len(mapStrAnimals))
	var animalIsExist bool
	for _, k := range keysAnimals {
		if k == "cat" {
			animalIsExist = true
		}
	}
	assert.True(t, animalIsExist)

	//************************
	//сбор всех ключей типа float64 в слайс
	mapFloatOne := map[float64]string{21.43: "longitude", 45.46: "latitude", 73.213: "longitude", 32.324: "latitude"}
	keysLongLatitude := GetKeysFromMap[float64, string](mapFloatOne)
	assert.Equal(t, len(keysLongLatitude), len(mapFloatOne))

	//************************
	//сложение любых числовых типов
	listIntForSum := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultInt = MySummation[int](listIntForSum)
	assert.Equal(t, resultInt, 45)
	listInt64ForSum := []int64{1, 2, 3, 4, -5, 6, 7, 8, 9}
	resultInt64 := MySummation[int64](listInt64ForSum)
	assert.Equal(t, resultInt64, int64(35))
	listUint64ForSum := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	resultUint64 := MySummation[uint64](listUint64ForSum)
	assert.Equal(t, resultUint64, uint64(45))
}
