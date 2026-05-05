// Мониторинг файлов
package filemonitoring_test

import (
	"context"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
)

// как отслеживать изменение файлов
// для этого будем использовать пакет fsnotify

func TestFsnotify(t *testing.T) {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	// создаем нового наблюдателя
	watcher, err := fsnotify.NewWatcher()
	// не забываем про обработку ошибок
	if err != nil {
		log.Fatal(err)
	}
	// не забываем удалить наблюдателя если он уже не нужен
	defer watcher.Close()

	// добавляем директорию которую будем наблюдать
	err = watcher.Add("../test/")
	if err != nil {
		log.Fatal(err)
	}

	// пишем гороутину для обработки событий от наблюдателя
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case event := <-watcher.Events:
				log.Printf("файл %s изменился, действие %s\n", event.Name, event.Op)

			case err := <-watcher.Errors:
				log.Println("принята ошибка:", err)

			}
		}
	}()

	// небольшая задержка что бы гороутина обработки событий гарантированно успела стартануть
	time.Sleep(1 * time.Second)

	// путь к тестовому файлу
	testFilePath := "../test/filetest.txt"

	// создаем новый файл
	f, err := os.OpenFile(testFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)

	// что то пишем в файл
	_, err = f.WriteString("writing any string")
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)

	// что то читаем из файла
	b := []byte{}
	_, err = f.Read(b)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(1 * time.Second)

	// изменяем права доступа к файлу
	if err := f.Chmod(0777); err != nil {
		log.Fatalln(err)
	}

	time.Sleep(3 * time.Second)

	// закрываем дескриптор файла
	f.Close()

	// удаляем файл
	if err = os.Remove(testFilePath); err != nil {
		log.Fatal(err)
	}

	cancel()
}
