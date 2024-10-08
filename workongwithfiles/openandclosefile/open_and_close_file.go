// - открытие и закрытие файла с различными атрибутами и правами доступа
package openandclosefile

import (
	"io/fs"
	"os"
)

// OpenFile открытие файла с различными атрибутами и правами доступа
// Можно использовать атрибуты как индивидуально так и комбинировать их.
// Например, O_CREATE|os.O_APPEND.
// Перечень атрибутов:
// os.O_RDONLY // Read only
// os.O_WRONLY // Write only
// os.O_RDWR // Read and write
// os.O_APPEND // Append to end of file
// os.O_CREATE // Create is none exist
// os.O_TRUNC // Truncate file when opening
// Перечень прав доступа и их аналог в Unix:
// ----------	0000	no permissions
// -rwx------	0700	read, write, & execute only for owner
// -rwxrwx---	0770	read, write, & execute for owner and group
// -rwxrwxrwx	0777	read, write, & execute for owner, group and others
// ---x--x--x	0111	execute
// --w--w--w-	0222	write
// --wx-wx-wx	0333	write & execute
// -r--r--r--	0444	read
// -r-xr-xr-x	0555	read & execute
// -rw-rw-rw-	0666	read & write
// -rwxr-----	0740	owner can read, write, & execute; group can only read; others have no permissions
func OpenFile(fileName string, attr int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(fileName, attr, perm)
}

// CloseFile закрывает файл
func CloseFile(f *os.File) {
	f.Close()
}
