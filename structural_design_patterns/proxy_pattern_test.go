package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxyExample(t *testing.T) {
	var object *VirtualProxy = NewVirtualProxy()
	result, err := object.performAction(2, 10)

	assert.Nil(t, err)
	assert.Equal(t, 5, result)
}
