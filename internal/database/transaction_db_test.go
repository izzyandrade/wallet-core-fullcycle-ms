package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TransactionDBTestSuite struct {
	suite.Suite
	transactionDB *TransactionDB
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", "file::memory:")
	s.Nil(err)

	s.db = db

	s.transactionDB = NewTransactionDB(s.db)

	db.Exec("CREATE TABLE clients (id VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, created_at DATETIME NOT NULL)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255) NOT NULL, client_id VARCHAR(255) NOT NULL, balance FLOAT NOT NULL, created_at DATETIME NOT NULL)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255) NOT NULL, account_from VARCHAR(255) NOT NULL, account_to VARCHAR(255) NOT NULL, amount FLOAT NOT NULL, created_at DATETIME NOT NULL, status VARCHAR(255) NOT NULL)")

	s.client1, _ = entity.NewClient("John Doe", "john.doe@example.com")
	s.client2, _ = entity.NewClient("Jane Doe", "jane.doe@example.com")

	s.accountFrom, _ = entity.NewAccount(s.client1)
	s.accountTo, _ = entity.NewAccount(s.client2)

	s.accountFrom.Balance = 1000
	s.accountTo.Balance = 1000

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreateTransaction() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)

	err = s.transactionDB.Create(transaction)
	s.Nil(err)
}
