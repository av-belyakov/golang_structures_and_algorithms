// Примеры обхода директорий
package traversingdirectories_test

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Walk не переходит по символическим ссылкам.
//
//Walk менее эффективен, чем WalkDir, представленный в Go 1.16, который позволяет избежать вызова os.Lstat
//для каждого посещенного файла или каталога.

type fileInfo struct {
	name  string
	size  int64
	isDir bool
}

func TestWalrDir(t *testing.T) {
	list := map[string]fileInfo{}

	err := filepath.WalkDir("../", func(path string, d fs.DirEntry, err error) error {
		if err != nil && err != fs.ErrNotExist {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		list[path] = fileInfo{
			name:  info.Name(),
			size:  info.Size(),
			isDir: d.IsDir(),
		}

		return nil
	})

	assert.NoError(t, err)

	for k, v := range list {
		fmt.Printf("path:'%s'\n\tisDir:'%t'\n\tname:'%s'\n\tsize:'%d'\n", k, v.isDir, v.name, v.size)
	}

	assert.NotEqual(t, len(list), 0)
}

func TestCalulateDirSize(t *testing.T) {
	var (
		size     int64
		countDir int
	)

	filepath.Walk("../", func(path string, info fs.FileInfo, err error) error {
		if err != nil && err != fs.ErrNotExist {
			return err
		}

		if info.IsDir() {
			countDir += 1
		}

		size += info.Size()

		return nil
	})

	t.Log("Direcotory size:", size)
	t.Log("Count directory:", countDir)

	assert.NotEqual(t, size, 0)
}
