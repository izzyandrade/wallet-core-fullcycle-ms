package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "john.doe@example.com")
	client2, _ := NewClient("Jane Doe", "jane.doe@example.com")
	account1, _ := NewAccount(client1)
	account2, _ := NewAccount(client2)

	account1.Deposit(1000.0)

	transaction, err := NewTransaction(account1, account2, 1000.0)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 1000.0, transaction.Amount)
	assert.Equal(t, account1, transaction.AccountFrom)
	assert.Equal(t, account2, transaction.AccountTo)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "john.doe@example.com")
	client2, _ := NewClient("Jane Doe", "jane.doe@example.com")
	account1, _ := NewAccount(client1)
	account2, _ := NewAccount(client2)

	account1.Deposit(1000.0)

	transaction, err := NewTransaction(account1, account2, 2000.0)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient funds", err.Error())
}

func TestCreateTransactionWithNegativeAmount(t *testing.T) {
	client1, _ := NewClient("John Doe", "john.doe@example.com")
	client2, _ := NewClient("Jane Doe", "jane.doe@example.com")
	account1, _ := NewAccount(client1)
	account2, _ := NewAccount(client2)

	account1.Deposit(1000.0)

	transaction, err := NewTransaction(account1, account2, -1000.0)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "amount must be positive", err.Error())
}

func TestCreateTransactionWithNilAccount(t *testing.T) {
	client, _ := NewClient("Jane Doe", "jane.doe@example.com")
	account, _ := NewAccount(client)

	transaction, err := NewTransaction(nil, account, 1000.0)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "account from and account to are required", err.Error())
}
