package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapterExample(t *testing.T) {
	assert.Equal(t, "10", AdapterExample(10))

	//для проверки на ошибку
	//assert.Nil(t, err)
}
