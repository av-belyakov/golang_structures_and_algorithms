package usebufferedreader

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseBufferedReader(t *testing.T) {
	var newFile string = "test_new_file.txt"

	err := os.WriteFile(newFile, []byte("Hi Men!\nThis is test file created on golang.\nWe just flushed, so there are no changes to revert.\nThe writer that you pass as an argument."), 0666)
	assert.NoError(t, err)

	f, err := os.OpenFile(newFile, os.O_RDONLY, 0666)
	assert.NoError(t, err)

	bufferedReader := bufio.NewReader(f)

	// Получаем байты без перемещения указателя
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(5)
	assert.NoError(t, err)

	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// Считывание и перемещение указателя
	numBytesRead, err := bufferedReader.Read(byteSlice)
	assert.NoError(t, err)

	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// Готов 1 байт. Ошибка, если нет байта для чтения
	myByte, err := bufferedReader.ReadByte()
	assert.NoError(t, err)

	fmt.Printf("Read 1 byte: %c\n", myByte)

	// Считывается до разделителя включительно
	// Возвращает фрагмент байта
	dataBytes, err := bufferedReader.ReadBytes('\n')
	assert.NoError(t, err)

	fmt.Printf("Read bytes: %s\n", dataBytes)

	// Читать до разделителя включительно
	// Возвращает строку
	dataString, err := bufferedReader.ReadString('\n')
	assert.NoError(t, err)

	fmt.Printf("Read string: %s\n", dataString)

	f.Close()
	err = os.Remove(newFile)
	assert.NoError(t, err)
}
