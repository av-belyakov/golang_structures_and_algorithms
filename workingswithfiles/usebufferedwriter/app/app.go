package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func App() {
	// как использовать bufio для записи в файл

	var fileName string = "test_file.txt"

	// открываем файл для записи
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// создаём новый писателя в буфер
	bufferedWriter := bufio.NewWriter(f)

	// записываем в буфер какие то байты
	bytesWritten, err := bufferedWriter.Write([]byte("this is FIRST test string\n"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("первая запись, записано", bytesWritten, "байт")

	// записываем в буфер какую то строку
	bytesWritten, err = bufferedWriter.WriteString("this is SECOND test string\n")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("вторая запись, записано", bytesWritten, "байт")

	// проверяем сколько данных хранится в буфере ожидания
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Println("в буфере хранится", unflushedBufferSize, "байт")

	// смотрим какой объем буфера доступен (4096 байт размер буфера без дополнительных настроек)
	bytesAvailable := bufferedWriter.Available()
	log.Println("всего доступно", bytesAvailable, "байт")

	// пишем данные из буфера в файл
	bufferedWriter.Flush()

	// проверяем наличие файла
	fileInfo, err := os.Stat(fileName)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	if (err == nil || !os.IsNotExist(err)) && strings.EqualFold(fileInfo.Name(), fileName) {
		log.Printf("OK, файл %s существует!\n", fileName)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("файл %s содержит следующие данные:\n%s\n", fileName, string(data))

	// закрываем файл
	f.Close()

	// удаляем файл
	if err = os.Remove(fileName); err != nil {
		log.Fatal(err)
	}

	//Отменить все изменения, внесенные в буфер, которые имеют
	// еще не записано в файл с помощью функции Flush()
	//bufferedWriter.Reset(bufferedWriter)

}
