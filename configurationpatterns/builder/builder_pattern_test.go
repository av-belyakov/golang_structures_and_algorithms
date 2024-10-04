package configurationpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilderPattern(t *testing.T) {
	builder := ConfigBuilder{}
	builder.Port(8080)
	cfg, err := builder.Build()
	assert.NoError(t, err)

	server, err := NewServerBuildPattern("localhost", cfg)
	assert.NoError(t, err)
	assert.Equal(t, "localhost:8080", server.Addr)
}
