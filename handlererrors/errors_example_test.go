package handlererrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerError(t *testing.T) {
	var errFoo error

	foo := func() error {
		errFoo = errors.New("generate new error for function Foo")

		return fmt.Errorf("func 'foo', error: %w", errFoo)
	}

	too := func() error {
		errFoo = errors.New("generate new error")

		return fmt.Errorf("func 'foo', error: %v", errFoo)
	}

	err := foo()
	assert.Error(t, err)
	assert.True(t, errors.Is(err, errFoo))

	err = too()
	assert.Error(t, err)
	assert.False(t, errors.Is(err, errFoo))
}
