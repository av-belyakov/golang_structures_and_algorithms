package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// как отслеживать изменение файлов
	// для этого будем использовать пакет fsnotify

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// создаем нового наблюдателя
	watcher, err := fsnotify.NewWatcher()
	// не забываем про обработку ошибок
	if err != nil {
		log.Fatal(err)
	}

	// незабываем удалить наблюдателя когда он не нужен
	defer watcher.Close()

	// добавляем директорию которую будем наблюдать
	if err = watcher.Add("./test/"); err != nil {
		log.Fatal(err)
	}

	// пишем гороутину для обработки события от наблюдателя
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case event := <-watcher.Events:
				log.Printf("файл %s изменился, действие %s\n", event.Name, event.Op.String())

			case err := <-watcher.Errors:
				log.Println("ошибка", err)
			}
		}
	}()

	// делаем небольшую задержку что бы гороутина обработчик событий успела стартануть
	time.Sleep(1 * time.Second)

	// путь к тестовому файлу
	testFilePath := "./test/filetest.txt"

	// создаем новый файл
	f, err := os.OpenFile(testFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	// что то пишем в файл
	if _, err := f.WriteString("writing any string"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	// изменяем права доступа к файлу
	if err = f.Chmod(0777); err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)

	// закрываем дескриптор файла
	f.Close()

	// удаляем файл
	if err = os.Remove(testFilePath); err != nil {
		log.Fatal(err)
	}

	time.Sleep(3 * time.Second)

	cancel()
}
