// Именованный конвеер
package navigatingnamedpipes_test

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/unix"
)

const pipeNamePath string = "name_pipe"

func namedPipeExists(pipePath string) (bool, error) {
	_, err := os.Stat(pipePath)
	if err == nil {
		//имя конвеера (pipe) существует
		return true, nil
	}
	if os.IsNotExist(err) {
		//имя конвеера (pipe) не существует
		return false, err
	}

	return false, err
}

func sendTask(pipe *os.File, data string) error {
	_, err := pipe.WriteString(data)
	if err != nil {
		return fmt.Errorf("error writing to named pipe: %v", err)
	}

	return nil
}

func readTask(pipe *os.File) error {
	fmt.Println("func 'readTask' start...")

	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		task := scanner.Text()
		fmt.Printf("Processing task: %s\n", task)
		if task == "EOD" {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading tasks: %w", err)
	}

	fmt.Println("func 'readTask' stop...")

	return nil
}

func TestMkfifo(t *testing.T) {
	isExist, _ := namedPipeExists(pipeNamePath)
	if !isExist {
		//если нет конвеера с таким именем, создаем
		err := unix.Mkfifo(pipeNamePath, 0666)
		assert.NoError(t, err)
	}

	// После создания файла "name_pipe" с помощью os.O_RDWR, который будет использоватся
	// для открытия канала для чтения. При этом отправляемые данные будут считыватся из канала
	pipeName, err := os.OpenFile(pipeNamePath, os.O_RDWR, os.ModeNamedPipe)
	assert.NoError(t, err)
	defer pipeName.Close()

	// Основная логика заключается в том, что одна программа отправляет задачи по конвейеру,
	// в то время как другая их считывает. Как только мы используем сканер, мы прекращаем
	// чтение новых задач, когда отправитель отправляет экземпляр строки "EOD".
	var wg sync.WaitGroup

	go func() {
		defer wg.Done()
		wg.Add(1)

		err := readTask(pipeName)
		assert.NoError(t, err)
	}()

	go func() {
		defer wg.Done()
		wg.Add(1)

		for i := range 10 {
			sendTask(pipeName, fmt.Sprintf("Task %d\n", i))
		}

		// закрываем pipeName
		sendTask(pipeName, "EOD\n")

		fmt.Println("All tasks sent.")
	}()

	wg.Wait()

	assert.True(t, true)
}
