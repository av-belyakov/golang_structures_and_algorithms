package readwithscanner

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadWithScanner(t *testing.T) {
	var newFile string = "test_new_file.txt"

	err := os.WriteFile(newFile, []byte("Hi Men!\nThis is test file created on golang.\nWe just flushed, so there are no changes to revert.\nThe writer that you pass as an argument."), 0666)
	assert.NoError(t, err)

	f, err := os.OpenFile(newFile, os.O_RDONLY, 0666)
	assert.NoError(t, err)

	scanner := bufio.NewScanner(f)

	// Сканер по умолчанию - bufio.ScanLines. Позволяет использовать сканирующие слова.
	// Также можно использовать пользовательскую функцию типа SplitFunc
	scanner.Split(bufio.ScanWords)

	// Сканировать в поисках следующего токена.
	success := scanner.Scan()
	if success == false {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err == nil {
			assert.Nil(t, err)

			log.Println("Scan completed and reached EOF")
		} else {
			assert.NoError(t, err)
		}
	}

	f.Close()
	err = os.Remove(newFile)
	assert.NoError(t, err)
}
