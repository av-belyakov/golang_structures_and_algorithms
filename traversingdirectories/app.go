package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

const Path_Dir = "../algorithms/"

func main() {
	//
	// обход директории и подсчет файлов и директорий
	//

	if err := WalkDir(); err != nil {
		fmt.Println(err)
	}
}

func WalkDir() error {
	var (
		sizeFiles            int64
		countDir, countFiles int
	)

	if err := filepath.WalkDir(Path_Dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil && err != fs.ErrNotExist {
			return err
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		if info.IsDir() {
			countDir++
		} else {
			countFiles++
			sizeFiles += info.Size()
		}

		fmt.Printf("\nпуть:%s\n", path)
		if info.IsDir() {
			fmt.Printf("  директория:%s\n", info.Name())
		} else {
			fmt.Printf("  файл:%s\n", info.Name())
		}

		fmt.Printf("\tразмер:%d byte\n", info.Size())

		return nil
	}); err != nil {
		return err
	}

	fmt.Println("-------------------------")
	fmt.Println("Всего директорий:", countDir)
	fmt.Println("Всего файлов:", countFiles)
	fmt.Println("Общий размер файлов:", sizeFiles)

	return nil
}
