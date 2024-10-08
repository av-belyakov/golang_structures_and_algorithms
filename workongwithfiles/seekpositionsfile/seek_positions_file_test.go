package seekpositionsfile

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeekPositionsFile(t *testing.T) {
	var newFile string = "test_new_file.txt"

	err := os.WriteFile(newFile, []byte("Hi Men!\nThis is test file created on golang.\n"), 0666)
	assert.NoError(t, err)

	f, err := os.OpenFile(newFile, os.O_RDONLY, 0666)
	assert.NoError(t, err)

	// Смещение - это количество байт для перемещения
	// Смещение может быть положительным или отрицательным
	var offset int64 = 5

	// Откуда берется точка отсчета для смещения
	// 0 = начальная позиция
	// 1 = текущая позиция
	// 2 = конец файла
	var whence int = 0
	newPosition, err := f.Seek(offset, whence)
	assert.NoError(t, err)

	fmt.Println("------", newPosition)

	f.Close()
	err = os.Remove(newFile)
	assert.NoError(t, err)
}
