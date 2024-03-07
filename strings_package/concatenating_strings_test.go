package stringspackage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strTest = "-I proto: Опция -I или --proto_path указывает путь к корневой директории с .proto файлами. Это нужно для того, чтобы компилятор смог найти импорты, если они есть. В нашем случае это директория proto"

func TestStringsConcat(t *testing.T) {
	strTest1 := ""
	testList := createTestStrList(strTest, 1000)

	for _, v := range testList {
		strTest1 += v
	}

	strTest2 := StringsConcat(testList)

	assert.Equal(t, len(strTest1), len(strTest2))
}

func BenchmarkSimpleStringsConcat(b *testing.B) {
	var listTest []string
	b.Run("Create test list", func(b *testing.B) {
		listTest = createTestStrList(strTest, 3000)
	})

	b.ResetTimer()

	b.Run("Simple concat", func(b *testing.B) {
		var strTest1 string
		for _, v := range listTest {
			strTest1 += v
		}
	})

	b.Run("Concat with strings.Builder{}", func(b *testing.B) {
		_ = StringsConcat(listTest)
	})

	b.Run("Concat with strings.Builder{} and Grow() method", func(b *testing.B) {
		_ = StringsConcatGrow(listTest)
	})
}

func createTestStrList(str string, num int) []string {
	list := make([]string, 0, num)

	for i := 0; i < num; i++ {
		list = append(list, str)
	}

	return list
}
