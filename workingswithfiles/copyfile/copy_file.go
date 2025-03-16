// - копирование байтов из исходного файла в целевой
package copyfile

import "io"

// CopyFile копирование байтов из исходного файла в целевой
func CopyFile(originalFile io.Reader, newFile io.Writer) (int64, error) {
	return io.Copy(newFile, originalFile)
}
