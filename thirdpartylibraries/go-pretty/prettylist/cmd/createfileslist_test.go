package prettylist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFilesList(t *testing.T) {
	rootPath, err := GetRootPath(ROOT_DIR)
	assert.NoError(t, err)
	fmt.Println("ROOT PATH:", rootPath)

	countFiles, countFolders, newList, err := CreateFilesAndFoldersList(rootPath)
	assert.NoError(t, err)
	assert.Greater(t, countFiles, 0)
	assert.Greater(t, countFolders, 0)

	fmt.Println("countFiles =", countFiles)
	fmt.Println("countFolders =", countFolders)
	fmt.Println("List:", newList)

	assert.True(t, false)
}
