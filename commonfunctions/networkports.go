package commonfunctions

import "math/rand"

// RandomPort вспомогательная функция генерации произвольного порта
func RandomPort() int {
	return rand.Intn(65535)
}

// DefaultHTTPPort функция возвращающая стандартный порт HTTP
func DefaultHTTPPort() int {
	return 80
}
