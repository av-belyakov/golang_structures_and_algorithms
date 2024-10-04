// Использование errgroup обеспечивает синхронизацию для группы подпрограмм и предоставляет решение для работы с ошибками и общими контекстами.
package errorshandling

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerErrgroup(t *testing.T) {
	listCircle := []Circle{
		{domain: "http://www.golang.org/"},
		{domain: "http://www.google.com/"},
		{domain: "http://www.somestupidname.com/"},
		//{domain: "http://my.example.net/"},
	}

	// result, err := handler(context.Background(), listCircle)
	result, err := ErrorGroup(listCircle)
	fmt.Println(result)

	//если при выполнение функции handler возникнет ошибка, например, из-за ошибок подключения
	// к несуществующему доменному имени то будет прекращена обработка всех выполняемых
	//горутин

	assert.NoError(t, err)
	assert.Equal(t, len(result), 3)
}

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
