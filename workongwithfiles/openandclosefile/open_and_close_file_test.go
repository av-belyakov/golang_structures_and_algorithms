package openandclosefile

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOpenAndCloseFile(t *testing.T) {
	var fileName string = "test_file.txt"

	f, err := OpenFile(fileName, os.O_APPEND|os.O_CREATE, fs.ModePerm)
	assert.NoError(t, err)

	CloseFile(f)

	err = os.Remove(fileName)
	assert.NoError(t, err)
}
