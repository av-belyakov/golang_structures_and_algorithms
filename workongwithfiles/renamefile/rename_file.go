// - переименование файла
package renamefile

import "os"

// RenameFile переименование файла
func RenameFile(oldName, newName string) error {
	return os.Rename(oldName, newName)
}
