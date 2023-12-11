package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdapterExample(t *testing.T) {
	var processor IAProcess = Adapter{adaptee: Adaptee{10}}

	assert.Equal(t, "10", processor.process())

	//для проверки на ошибку
	//assert.Nil(t, err)
}
