package usebufferedwriter

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUseBufferedWriter(t *testing.T) {
	var fileName string = "test_file_name.txt"

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	assert.NoError(t, err)

	bufferedWriter := bufio.NewWriter(f)

	bytesWritten, err := bufferedWriter.Write([]byte{65, 66, 67, 01, 23, 34, 134})
	assert.NoError(t, err)
	assert.Greater(t, bytesWritten, 3)

	bytesWritten, err = bufferedWriter.WriteString("Buffered string\n")
	assert.NoError(t, err)
	assert.Greater(t, bytesWritten, 10)

	// проверяем сколько данных хранится в буфере ожидания
	unflushedBufferSize := bufferedWriter.Buffered()
	assert.Greater(t, unflushedBufferSize, 15)

	//смотрим какой объем буфера доступен
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Записать буфер памяти на диск
	bufferedWriter.Flush()

	//Отменить все изменения, внесенные в буфер, которые имеют
	// еще не записано в файл с помощью функции Flush()
	bufferedWriter.Reset(bufferedWriter)

	f.Close()
	err = os.Remove(fileName)
	assert.NoError(t, err)
}
