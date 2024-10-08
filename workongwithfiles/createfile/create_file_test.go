package createfile

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSimpleFile(t *testing.T) {
	var fileName string = "test_simplefile.txt"

	f, err := SimpleCreateFile(fileName)
	assert.NoError(t, err)
	f.Close()

	fileInfo, err := os.Stat(fileName)
	assert.True(t, !os.IsNotExist(err))

	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("Size in bytes:", fileInfo.Size())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Printf("System interface type: %T\n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())

	err = os.Remove(fileName)
	assert.NoError(t, err)
}

func TestCreateTruncateFile(t *testing.T) {
	var (
		fileName string = "test_truncatefile.txt"
		fsize    int64  = 100
	)

	f, err := TruncateCreateFile(fileName, fsize)
	f.Close()
	assert.NoError(t, err)

	fileInfo, err := os.Stat(fileName)
	assert.True(t, !os.IsNotExist(err))
	assert.Equal(t, fileInfo.Size(), fsize)

	err = os.Remove(fileName)
	assert.NoError(t, err)
}
