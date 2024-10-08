// - установка конфигурационных опций на основе шаблона "Строитель" (Builder pattern)
package configurationpatterns

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/av-belyakov/golang_structures_and_algorithms/commonfunctions"
)

//Так как при использовании структуры не указанные поля инициализируются со значением "по
//умолчанию", то необходимо отличать установленное пользователем значение, например сетевого порта
//установленного в "0" от значения инициализированного в структуре "по умолчанию". Для этого
//можно использовать вспомогательную структуру.

// Config основная конфигурационноая структура
type Config struct {
	Port int
}

// ConfigBuilder вспомогательная конфигурационноая структура
type ConfigBuilder struct {
	port *int
}

// Port метод позволяющий задать конфигурацию порта
func (b *ConfigBuilder) Port(
	port int) *ConfigBuilder {
	b.port = &port
	return b
}

// Build метод строящий конфигурацию
func (b *ConfigBuilder) Build() (Config, error) {
	cfg := Config{}

	if b.port == nil {
		cfg.Port = commonfunctions.DefaultHTTPPort()
	} else {
		if *b.port == 0 {
			cfg.Port = commonfunctions.RandomPort()
		} else if *b.port < 0 {
			return Config{}, errors.New("port should be positive")
		} else {
			cfg.Port = *b.port
		}
	}
	return cfg, nil
}

//Сначала клиент создает конструктор конфигурации и использует его для настройки необязательного поля, такого как
//порт. Затем он вызывает метод сборки и проверяет наличие ошибок. Если все в порядке, настройка передается на
//новый сервер.
//При таком подходе создается обработчик управления портом. Не обязательно передавать указатель integer, поскольку
//метод Port принимает целое число. Однако нам все равно нужно передать структуру config, которая может быть пустой,
//если клиент хочет использовать конфигурацию по умолчанию:
//
//server, err := httplib.NewServer("localhost", nil)
//
//Другой недостаток, в некоторых ситуациях, связан с управлением ошибками. В языках программирования, где генерируются
//исключения, методы конструктора, такие как Port, могут вызывать исключения, если ввод неверен. Если мы хотим сохранить
//возможность цепочки вызовов, функция не может возвращать ошибку.

// NewServerBuildPattern функция генерирующая новый сервер
func NewServerBuildPattern(addr string, config Config) (*http.Server, error) {

	// ...

	return &http.Server{Addr: fmt.Sprintf("%s:%d", addr, config.Port)}, nil
}
