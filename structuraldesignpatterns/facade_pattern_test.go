package structuraldesignpatterns

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFacadeExample(t *testing.T) {
	var (
		customer = "Thomas Smith"
		account  = "Savings"

		srcAccountId = "21456"
		dstAccountId = "87345"
	)

	var facade = NewBranchManagerFacade()

	_, _ = facade.CreateCustomerAccount(customer, account)
	c, a := facade.GetCustomerAccount()

	fmt.Println("customer: ", c)
	fmt.Println("account: ", a)

	assert.Equal(t, customer, c.name)
	assert.Equal(t, account, a.accountType)

	_ = facade.CreateTransaction(srcAccountId, dstAccountId, 1000)

	assert.Equal(t, srcAccountId, facade.GetTransactionSrcAccountId())
	assert.Equal(t, dstAccountId, facade.GetTransactionDstAccountId())
}
