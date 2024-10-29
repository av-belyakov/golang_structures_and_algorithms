package prettylist

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Element struct {
	Name string
	Next []Element
}

func CreateFilesAndFoldersList(currentPath string) (int, int, []Element, error) {
	var countFiles, countFolders int
	finalyList := []Element(nil)

	list, err := os.ReadDir(currentPath)
	if err != nil {
		return countFiles, countFolders, finalyList, err
	}

	for _, v := range list {
		if info, err := v.Info(); err == nil {
			if info.IsDir() {
				countFolders++

				if cfiles, cfolders, elem, err := CreateFilesAndFoldersList(fmt.Sprintf("%s/%s", currentPath, v.Name())); err == nil {
					finalyList = append(finalyList, Element{
						Name: v.Name(),
						Next: elem,
					})

					countFiles += cfiles
					countFolders += cfolders
				}

				continue
			}

			finalyList = append(finalyList, Element{Name: v.Name()})
			countFiles++
		}
	}

	return countFiles, countFolders, finalyList, nil
}

func GetRootPath(rootDir string) (string, error) {
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return "", err
	}

	tmp := strings.Split(currentDir, "/")

	if tmp[len(tmp)-1] == rootDir {
		return currentDir, nil
	}

	var currentPath string = ""
	for _, v := range tmp {
		currentPath += v + "/"

		if v == rootDir {
			return currentPath, nil
		}
	}

	return currentPath, nil
}
