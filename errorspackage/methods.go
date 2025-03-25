package errorspackage

import "fmt"

func (e *CreateCaseError) Error() string {
	return fmt.Sprintf("%s:%v", e.Type, e.Err)
}

func (e *CreateCaseError) Is(target error) bool {
	ccerr, ok := target.(*CreateCaseError)
	if !ok {
		return false
	}

	return ccerr.Type == e.Type
}

/*
func (e *CreateCaseError) As(target any) bool {
	if _, ok := target.(CreateCaseError); ok {
		return true
	}

	return false
}
*/
