package fucnoptions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewServerWithFuncOpts(t *testing.T) {
	var (
		host string = "example.ru"
		port int    = 9090
	)

	server := NewServer(
		WithHost(host),
		WithNetPort(port),
		WithLogs(true),
		WithSSL(true),
		WithTimeout(10*time.Second))

	assert.Equal(t, server.Host, host)
	assert.Equal(t, server.Port, port)
	assert.True(t, server.EnableLogs)
	assert.True(t, server.WithSSL)
}
