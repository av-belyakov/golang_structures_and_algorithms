package supportfunctions

import (
	"fmt"
	"os"
	"slices"
)

// DeleteFilesInDirectory удаляет перечисленные файлы в заданной директории
func DeleteFilesInDirectory(dir string, files []string) error {
	dirInfo, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, v := range files {
		if slices.ContainsFunc(dirInfo, func(f os.DirEntry) bool {
			return v == f.Name()
		}) {
			filePath := fmt.Sprintf("%s/%s", dir, v)
			if err := os.RemoveAll(filePath); err != nil {
				return err
			}
		}
	}

	return nil
}
