package checkfileexist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChechFileExist(t *testing.T) {
	assert.True(t, CheckFileExist("check_file_exist.go"))
}
