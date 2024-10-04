package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoratorExample(t *testing.T) {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	assert.Equal(t, "hotdog", decorator.process("hotdog", 3))

	decorator.processInstance = process
	assert.Equal(t, "hotdoghotdoghotdoghotdoghotdog", decorator.process("hotdog", 5))
}
