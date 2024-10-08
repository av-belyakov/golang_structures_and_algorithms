package writebyteorstringtofile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteByteToFile(t *testing.T) {
	var fileName string = "test_file_name.txt"

	f, err := os.Create(fileName)
	assert.NoError(t, err)

	countByte, err := WriteByteToFile(f, []byte("Hi people!"))
	assert.NoError(t, err)
	assert.Greater(t, countByte, 5)

	f.Close()
	err = os.Remove(fileName)
	assert.NoError(t, err)
}
