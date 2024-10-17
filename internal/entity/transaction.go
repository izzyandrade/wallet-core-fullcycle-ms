package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type TransactionStatus string

const (
	TransactionPending   TransactionStatus = "pending"
	TransactionConfirmed TransactionStatus = "confirmed"
	TransactionCanceled  TransactionStatus = "canceled"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
	Status      TransactionStatus
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
		Status:      TransactionPending,
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	transaction.Execute()

	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil || t.AccountTo == nil {
		return errors.New("account from and account to are required")
	}
	if t.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}

func (t *Transaction) Execute() {
	t.AccountFrom.Withdraw(t.Amount)
	t.AccountTo.Deposit(t.Amount)
	t.Status = TransactionConfirmed
}

func (t *Transaction) Cancel() {
	if t.Status == TransactionConfirmed {
		return
	}
	t.Status = TransactionCanceled
}
