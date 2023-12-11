package structuraldesignpatterns

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrivateClassDataExample(t *testing.T) {
	var account *PCDAccount = &PCDAccount{CustomerName: "John Smith"}
	account.setDetails("4532", "current")

	//сокрытие данных из PCDAccount.details.id и PCDAccount.details.type для
	//метода обработки JSON
	jsonAccount, _ := json.Marshal(account)
	assert.Equal(t, "{\"CustomerName\":\"John Smith\"}", string(jsonAccount))

	//доступ к частным полям можно получить только специальными методами
	assert.Equal(t, "4532", account.getId())
	assert.Equal(t, "current", account.getAccountType())
}
