package writebyteorstringtofile

import "os"

// WriteByteToFile запись среза байт в файл
func WriteByteToFile(f *os.File, b []byte) (int, error) {
	return f.Write(b)
}
