// - установка конфигурационных опций на основе "шаблона функциональных опций" (Functional options pattern)
package fucnoptions

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/av-belyakov/golang_structures_and_algorithms/commonfunctions"
)

type options struct {
	port *int
}

type Option func(options *options) error

// Здесь With Port возвращает замыкание. Замыкание - это анонимная функция, которая ссылается
// на переменные извне своего тела; в данном случае на переменную port. Замыкание учитывает
// тип параметра и реализует логику проверки порта. Каждое поле конфигурации требует создания
// общедоступной функции (которая по соглашению начинается с префикса With), содержащей
// аналогичную логику: проверку входных данных при необходимости и обновление структуры конфигурации.
func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

// NewServerFunctionalOptionsPattern создает новый сервре с определенной конфигурацией
func NewServerFunctionalOptionsPattern(addr string, opts ...Option) (*http.Server, error) {
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return nil, err
		}
	}
	// At this stage, the options struct is built and contains the config
	// Therefore, we can implement our logic related to port configuration
	var port int
	if options.port == nil {
		port = commonfunctions.DefaultHTTPPort()
	} else {
		if *options.port == 0 {
			port = commonfunctions.RandomPort()
		} else {
			port = *options.port
		}
	}

	// ...

	return &http.Server{Addr: fmt.Sprintf("%s:%d", addr, port)}, nil
}
