package renamefile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRenameFile(t *testing.T) {
	var (
		oldName string = "test_file_1.txt"
		newName string = "test_file_2.txt"
	)

	f, err := os.Create(oldName)
	assert.NoError(t, err)
	f.Close()

	err = RenameFile(oldName, newName)
	assert.NoError(t, err)

	_, err = os.Stat(newName)
	assert.True(t, !os.IsNotExist(err))

	err = os.Remove(newName)
	assert.NoError(t, err)
}
