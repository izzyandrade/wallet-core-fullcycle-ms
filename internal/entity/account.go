package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        string
	Client    *Client
	Balance   float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(client *Client) (*Account, error) {
	account := &Account{
		ID:        uuid.New().String(),
		Client:    client,
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := account.Validate(); err != nil {
		return nil, err
	}

	return account, nil
}

func (a *Account) Validate() error {
	if a.Client == nil {
		return errors.New("client is required")
	}
	return nil
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	a.Balance += amount
	a.UpdatedAt = time.Now()
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be positive")
	}
	if amount > a.Balance {
		return errors.New("insufficient funds")
	}
	a.Balance -= amount
	a.UpdatedAt = time.Now()
	return nil
}
