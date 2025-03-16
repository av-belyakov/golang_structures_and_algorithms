package writebyteorstringtofile

import "os"

// WriteStringToFile запись строки в файл
func WriteStringToFile(f *os.File, str string) (int, error) {
	return f.WriteString(str)
}
