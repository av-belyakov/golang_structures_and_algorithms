package errorspackage_test

import (
	"encoding/json"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/golang_structures_and_algorithms/errorspackage"
)

func TestErrorAs(t *testing.T) {
	errFirst := &errorspackage.CreateCaseError{}
	errSecond := &errorspackage.CreateCaseError{}
	errThird := &errorspackage.CreateCaseError{}

	errFirst.Type = "json"
	errFirst.Err = &json.SyntaxError{}

	errSecond.Type = "json"
	errSecond.Err = &json.SyntaxError{}

	errThird.Type = "file"
	errThird.Err = os.ErrNotExist

	var e *errorspackage.CreateCaseError
	assert.True(t, errors.As(errFirst, &e))
	assert.True(t, errors.Is(errFirst, errFirst))
	assert.False(t, errors.Is(errThird, &errorspackage.CreateCaseError{Type: "json"}))
	assert.True(t, errors.Is(errThird, &errorspackage.CreateCaseError{Type: "file"}))
}
