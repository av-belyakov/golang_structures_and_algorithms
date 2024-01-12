package configurationpatterns

import (
	"errors"
	"fmt"
	"net/http"
)

//Установка конфигурационных опций на основе "шаблона функциональных опций" (Functional options pattern)

type options struct {
	port *int
}

type Option func(options *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return errors.New("port should be positive")
		}
		options.port = &port
		return nil
	}
}

//Здесь With Port возвращает замыкание. Замыкание - это анонимная функция, которая ссылается
//на переменные извне своего тела; в данном случае на переменную port. Замыкание учитывает
//тип параметра и реализует логику проверки порта. Каждое поле конфигурации требует создания
//общедоступной функции (которая по соглашению начинается с префикса With), содержащей
//аналогичную логику: проверку входных данных при необходимости и обновление структуры конфигурации.

// ********************
// Пример использования
// ********************
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
		port = defaultHTTPPort()
	} else {
		if *options.port == 0 {
			port = randomPort()
		} else {
			port = *options.port
		}
	}

	// ...

	return &http.Server{Addr: fmt.Sprintf("%s:%d", addr, port)}, nil
}
