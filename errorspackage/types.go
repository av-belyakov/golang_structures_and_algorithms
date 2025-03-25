// Некоторые примеры работ с ошибками
package errorspackage

type CreateCaseError struct {
	Type string
	Err  error
}
