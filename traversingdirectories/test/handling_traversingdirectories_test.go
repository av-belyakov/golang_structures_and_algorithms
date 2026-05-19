// Примеры обхода директорий
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

func TestWalkDir(t *testing.T) {
	list := map[string]fileInfo{}

	//err := filepath.WalkDir("../", func(path string, d fs.DirEntry, err error) error {
	err := filepath.WalkDir("../../algorithms/", func(path string, d fs.DirEntry, err error) error {
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
		fmt.Printf("\npath:%s\n", k)
		if v.isDir {
			fmt.Printf("\tdirectory:%s\n", v.name)
		} else {
			fmt.Printf("\tfile:%s\n", v.name)
		}

		fmt.Printf("\tsize:%d byte\n", v.size)
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

	t.Log("Всего директорий:", countDir)
	t.Log("Общий размер директорий:", size)

	assert.NotEqual(t, size, 0)
}
