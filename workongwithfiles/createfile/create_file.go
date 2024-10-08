// - создание файлов
package createfile

import "os"

// SimpleCreateFile простое создание файла
func SimpleCreateFile(name string) (*os.File, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// TruncateCreateFile создание файла ограниченного размера. Например, был создан файл размером 100.
// Если размер файла меньше 100 байт, исходное содержимое останется в начале, а оставшееся
// пространство будет заполнено нулевыми байтами. Если размер файла больше 100 байт. Все, что
// превышает 100 байт, будет потеряно.
func TruncateCreateFile(name string, size int64) (*os.File, error) {
	f, err := SimpleCreateFile(name)
	if err != nil {
		return nil, err
	}

	err = os.Truncate(name, size)

	return f, err
}
