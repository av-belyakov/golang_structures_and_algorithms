package copyfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	var (
		newFile      string = "test_new_file.txt"
		originalFile string = "test_original_file.txt"
	)

	err := os.WriteFile(newFile, []byte("Hi Men!\nThis is test file created on golang.\n"), 0666)
	assert.NoError(t, err)

	origf, err := os.OpenFile(newFile, os.O_RDONLY, 0666)
	assert.NoError(t, err)

	newf, err := os.OpenFile(originalFile, os.O_WRONLY|os.O_CREATE, 0666)
	assert.NoError(t, err)

	countByte, err := CopyFile(origf, newf)
	assert.NoError(t, err)
	assert.Greater(t, countByte, int64(10))

	origf.Close()
	newf.Close()

	err = os.Remove(newFile)
	assert.NoError(t, err)

	err = os.Remove(originalFile)
	assert.NoError(t, err)
}
