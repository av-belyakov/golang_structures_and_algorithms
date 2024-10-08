package compressuncompressfile

import (
	"compress/gzip"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressUncompressFile(t *testing.T) {
	var (
		zipped   string = "zippedfile.txt.gzip"
		unzipped string = "unzippedfile.txt"
	)

	//*** запись в архив
	gzipFile, err := os.OpenFile(zipped, os.O_WRONLY|os.O_CREATE, 0666)
	assert.NoError(t, err)

	gzipWriter := gzip.NewWriter(gzipFile)

	// Когда мы записываем данные в файл gzip writer он, в свою очередь, сжимает содержимое
	// и затем записывает его в базовый файл также для записи файлов. Нам не нужно
	// беспокоиться о том, как работает все сжатие, поскольку мы просто
	// используем его как простой интерфейс записи,  в который мы отправляем байты.
	_, err = gzipWriter.Write([]byte("Hi Men!\nThis is test file created on golang."))
	assert.NoError(t, err)

	gzipWriter.Close()
	gzipFile.Close()

	//*** чтение из архива
	gzipFile, err = os.OpenFile(zipped, os.O_RDONLY, 0666)
	assert.NoError(t, err)

	//выполняем чтение gzip файла
	gzipReader, err := gzip.NewReader(gzipFile)
	assert.NoError(t, err)

	//создаем файл в который будет разархивироватся файл gzip
	outfileWriter, err := os.OpenFile(unzipped, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	assert.NoError(t, err)

	//копируем содержимое файла gzip в обычный файл
	_, err = io.Copy(outfileWriter, gzipReader)
	assert.NoError(t, err)

	gzipReader.Close()
	outfileWriter.Close()

	err = os.Remove(zipped)
	assert.NoError(t, err)
	err = os.Remove(unzipped)
	assert.NoError(t, err)
}
