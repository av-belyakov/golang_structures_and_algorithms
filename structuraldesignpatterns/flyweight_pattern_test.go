package structuraldesignpatterns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlyweightExample(t *testing.T) {
	var factory = DataTransferObjectFactory{make(map[string]DataTransferObject)}

	var customer DataTransferObject = factory.getDataTransferObject("customer")
	assert.Equal(t, "1", customer.getId())

	var employee DataTransferObject = factory.getDataTransferObject("employee")
	assert.Equal(t, "2", employee.getId())

	var manager DataTransferObject = factory.getDataTransferObject("manager")
	assert.Equal(t, "3", manager.getId())

	var address DataTransferObject = factory.getDataTransferObject("address")
	assert.Equal(t, "4", address.getId())
}
