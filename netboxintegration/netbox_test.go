package netboxintegration_test

import (
	"fmt"
	"testing"

	"github.com/av-belyakov/golang_structures_and_algorithms/netboxintegration"
	"github.com/stretchr/testify/assert"
)

func TestGetPrefix(t *testing.T) {
	// токен доступа создаём в UI Netbox и добавляем различные разрешения
	// на действия группы
	nbClient, err := netboxintegration.NewClient("localhost", 8080, "HOBxYdR3Q5qGNVTtTVqkKiNrZChOaC9Z69QTJN6Q")
	if err != nil {
		t.Fatal(err)
	}

	// подробнее о запросах смотреть на https://netboxlabs.com/docs/netbox/integrations/rest-api/
	query := "/api/status/"

	res, size, err := nbClient.Get(t.Context(), query)
	assert.NoError(t, err)
	assert.Equal(t, 200, size)

	fmt.Println("Response:", string(res))
}
