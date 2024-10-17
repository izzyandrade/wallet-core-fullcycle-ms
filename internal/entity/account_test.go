package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, err := NewAccount(client)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, 0.0, account.Balance)
}

func TestNewAccountWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("", "")
	account, err := NewAccount(client)
	assert.NotNil(t, err)
	assert.Nil(t, account)
	assert.Equal(t, "client is required", err.Error())
}

func TestAccountDeposit(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, _ := NewAccount(client)
	err := account.Deposit(100.0)
	assert.Nil(t, err)
	assert.Equal(t, 100.0, account.Balance)
}

func TestAccountDepositWithNegativeAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, _ := NewAccount(client)
	err := account.Deposit(-100.0)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be positive", err.Error())
}

func TestAccountWithdraw(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, _ := NewAccount(client)
	err := account.Deposit(100.0)
	assert.Nil(t, err)
	err = account.Withdraw(50.0)
	assert.Nil(t, err)
	assert.Equal(t, 50.0, account.Balance)
}

func TestAccountWithdrawWithNegativeAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, _ := NewAccount(client)
	err := account.Deposit(100.0)
	assert.Nil(t, err)
	err = account.Withdraw(-50.0)
	assert.NotNil(t, err)
	assert.Equal(t, "amount must be positive", err.Error())
}

func TestAccountWithdrawWithInsufficientFunds(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, _ := NewAccount(client)
	err := account.Deposit(100.0)
	assert.Nil(t, err)
	err = account.Withdraw(200.0)
	assert.NotNil(t, err)
	assert.Equal(t, "insufficient funds", err.Error())
}
