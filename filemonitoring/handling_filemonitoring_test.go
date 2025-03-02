// Мониторинг файлов
package filemonitoring_test

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
)

func TestFsnotify(t *testing.T) {
	watchPath := "../filemonitoring/"

	watcher, err := fsnotify.NewWatcher()
	assert.NoError(t, err)

	err = watcher.Add(watchPath)
	assert.NoError(t, err)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				//прием событий
				fmt.Printf("Event: %s, Operation: %s\n", event.Name, event.Op)
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		pathTestFile := "./filetest.txt"

		// выполняем взаимодействия с файлами
		f, err := os.OpenFile(pathTestFile, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatalln(err)
		}
		defer func(f *os.File, fp string) {
			f.Close()
			os.Remove(fp)
		}(f, pathTestFile)

		_, err = f.WriteString("write any string")
		if err != nil {
			log.Fatalln(err)
		}

		b := []byte{}
		_, err = f.Read(b)
		if err != nil {
			log.Fatalln(err)
		}

		if f.Chmod(0777); err != nil {
			log.Fatalln(err)
		}
	}()

	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, os.Interrupt)

	<-chSignal

	t.Log("Test stop")
	watcher.Close()
}
