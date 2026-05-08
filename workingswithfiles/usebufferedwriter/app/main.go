package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	// как использовать bufio для записи данных в файл

	fileName := "test_file.txt"

	// открываем файл для записи
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// создаём нового писателя в буфер
	bufferWriter := bufio.NewWriter(f)

	// записываем в буфер какие то байты
	countWrited, err := bufferWriter.Write([]byte("This is FIRST test string\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("первая запись, записано", countWrited, "байт")

	// записываем в буфер некоторую строку
	countWrited, err = bufferWriter.WriteString("This is SECOND test string\n")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("вторая запись, записано", countWrited, "байт")

	// проверяем сколько данных хранится в буфере ожидания
	unflushedBufferSize := bufferWriter.Buffered()
	log.Println("в буфере хранится", unflushedBufferSize, "байт")

	// смотрим какой объем доступен (4096 байт обычный размер буфера без доп. параметров)
	bytesAvailable := bufferWriter.Available()
	log.Println("всего доступно", bytesAvailable, "байт")

	// пишем данные из буфера в файл
	bufferWriter.Flush()

	// проверяем наличие файла
	fileInfo, err := os.Stat(fileName)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	if (err == nil || !os.IsNotExist(err)) && strings.EqualFold(fileName, fileInfo.Name()) {
		log.Printf("OK, файл %s существует!", fileName)
	}

	// читаем из файла
	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("файл %s содержит следующие данные:\n%s\n", fileName, string(data))

	// не забываем закрыть файл
	f.Close()

	// удаляем файл
	if err = os.Remove(fileName); err != nil {
		log.Fatal(err)
	}
}
