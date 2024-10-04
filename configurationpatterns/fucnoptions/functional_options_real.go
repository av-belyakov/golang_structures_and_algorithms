package fucnoptions

import "time"

// server сетевой сервер
type server struct {
	WithSSL    bool
	EnableLogs bool
	Port       int
	Host       string
	Timeout    time.Duration
}

// serverOption функциональные опции для настройки сетевого сервера
type serverOption func(*server)

// WithHost задает доменное имя или ip
func WithHost(host string) serverOption {
	return func(s *server) {
		s.Host = host
	}
}

// WithNetPort функция задающая сетевой порт
func WithNetPort(port int) serverOption {
	return func(s *server) {
		s.Port = port
	}
}

// WithTimeout функция задающая какой либо таймаут
func WithTimeout(timeout time.Duration) serverOption {
	return func(s *server) {
		s.Timeout = timeout
	}
}

// WithLogs функция включающая или выключающая логирование
func WithLogs(enabled bool) serverOption {
	return func(s *server) {
		s.EnableLogs = enabled
	}
}

// WithSSL функция включающая или выключающая SSL соединение
func WithSSL(enabled bool) serverOption {
	return func(s *server) {
		s.WithSSL = enabled
	}
}

// NewServer конструктор генерирующий новый сервер
func NewServer(opts ...serverOption) *server {
	server := &server{
		Port:       8080,
		Timeout:    60,
		EnableLogs: false,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}
