package checkrwpermissions

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIsPermissions(t *testing.T) {
	fileName := "our_test_file.txt"

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	f.Close()
	assert.False(t, CheckIsPermissions(err))

	err = os.Remove(fileName)
	assert.NoError(t, err)

	f, err = os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0666)
	f.Close()
	assert.False(t, CheckIsPermissions(err))

	err = os.Remove(fileName)
	assert.NoError(t, err)
}
