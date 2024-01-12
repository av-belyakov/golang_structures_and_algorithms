package configurationpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunctionalOptionsPattern(t *testing.T) {
	//с заданными параметрами
	server, err := NewServerFunctionalOptionsPattern("localhost", WithPort(8080))
	assert.NoError(t, err)
	assert.Equal(t, "localhost:8080", server.Addr)

	//со значениями по умолчанию (по умолчанию ставится порт 80)
	server, err = NewServerFunctionalOptionsPattern("localhost")
	assert.NoError(t, err)
	assert.Equal(t, "localhost:80", server.Addr)
}
