// - проверка существования файлов
package checkfileexist

import "os"

// CheckFileExist проверяет наличие файла
func CheckFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
}
