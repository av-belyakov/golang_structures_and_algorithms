// Шаблон "Данные частного класса" (Private class data)
package structuraldesignpatterns

import (
	"encoding/json"
	"fmt"
)

// Шаблон данных частного класса защищает данные внутри класса. Этот шаблон инкапсулирует
//инициализацию данных класса. Права на запись свойств в частном классе
//защищены, и свойства устанавливаются во время построения. Шаблон частного класса печатает
//доступ к информации, защищая ее в классе, который остается в состоянии. Инкапсуляция
//инициализации данных класса - это сценарий, в котором применим этот шаблон.

// AccountDetails struct
type AccountDetails struct {
	id          string
	accountType string
}

// Account struct
type PCDAccount struct {
	details      *AccountDetails
	CustomerName string
}

// Account class method setDetails
func (account *PCDAccount) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

// Account class method getId
func (account *PCDAccount) getId() string {
	return account.details.id
}

// Account class method getAccountType
func (account *PCDAccount) getAccountType() string {
	return account.details.accountType
}

func PrivateClassDataExample() {
	var account *PCDAccount = &PCDAccount{CustomerName: "John Smith"}
	account.setDetails("4532", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))
	fmt.Println("Account Id", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
