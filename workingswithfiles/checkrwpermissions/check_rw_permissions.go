// - проверка прав доступа файла на запись и чтение
package checkrwpermissions

import "os"

// CheckIsPermissions проверка прав доступа файла на запись и чтение
func CheckIsPermissions(err error) bool {
	return os.IsPermission(err)
}
