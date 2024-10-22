package create_transaction

import (
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}
type AccountGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	return m.Called(transaction).Error(0)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john.doe@example.com")
	account1, _ := entity.NewAccount(client1)
	account1.Deposit(1000.0)

	client2, _ := entity.NewClient("Jane Doe", "jane.doe@example.com")
	account2, _ := entity.NewAccount(client2)
	account2.Deposit(1000.0)

	accountMock := &AccountGatewayMock{}
	accountMock.On("FindByID", account1.ID).Return(account1, nil)
	accountMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionMock := &TransactionGatewayMock{}
	transactionMock.On("Create", mock.Anything).Return(nil)

	input := CreateTransactionInputDTO{
		AccountIDFrom: account1.ID,
		AccountIDTo:   account2.ID,
		Amount:        100.0,
	}

	uc := NewCreateTransactionUseCase(transactionMock, accountMock)

	output, err := uc.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)

	accountMock.AssertExpectations(t)
	transactionMock.AssertExpectations(t)

	accountMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionMock.AssertNumberOfCalls(t, "Create", 1)

	assert.Equal(t, account1.Balance, 900.0)
	assert.Equal(t, account2.Balance, 1100.0)
}
