package database

import (
	"database/sql"
	"testing"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", "file::memory:")
	s.Nil(err)

	s.db = db
	s.accountDB = NewAccountDB(db)
	s.client, err = entity.NewClient("John Doe", "john.doe@example.com")
	s.Nil(err)

	db.Exec("CREATE TABLE clients (id VARCHAR(255) NOT NULL, name VARCHAR(255) NOT NULL, email VARCHAR(255) NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, PRIMARY KEY (id))")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255) NOT NULL, client_id VARCHAR(255) NOT NULL, balance FLOAT NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NOT NULL, PRIMARY KEY (id), FOREIGN KEY (client_id) REFERENCES clients (id))")
}

func (s *AccountDBTestSuite) TearDownSuite() {
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Close()
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestSave() {
	account, err := entity.NewAccount(s.client)
	s.Nil(err)

	errSave := s.accountDB.Save(account)
	s.Nil(errSave)

	s.NotNil(account.ID)
}

func (s *AccountDBTestSuite) TestFindByID() {

	stmt, err := s.db.Prepare("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	s.Nil(err)
	defer stmt.Close()

	_, err = stmt.Exec(s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, s.client.UpdatedAt)
	s.Nil(err)

	stmt, err = s.db.Prepare("SELECT * FROM clients WHERE id = ?")
	s.Nil(err)

	foundClient := &entity.Client{}
	row := stmt.QueryRow(s.client.ID)
	err = row.Scan(&foundClient.ID, &foundClient.Name, &foundClient.Email, &foundClient.CreatedAt, &foundClient.UpdatedAt)
	s.Nil(err)

	account, err := entity.NewAccount(foundClient)
	s.Nil(err)

	errSave := s.accountDB.Save(account)
	s.Nil(errSave)

	accountFound, err := s.accountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountFound.ID)
	s.Equal(account.Client.ID, accountFound.Client.ID)
}
