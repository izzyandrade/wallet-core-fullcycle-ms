package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/entity"
)

type AccountDB struct {
	DB *sql.DB
}

func NewAccountDB(db *sql.DB) *AccountDB {
	return &AccountDB{DB: db}
}

func (a *AccountDB) FindByID(id string) (*entity.Account, error) {
	var account entity.Account
	var client entity.Client

	account.Client = &client

	fmt.Printf("Finding account with ID: %s\n", id)

	stmt, err := a.DB.Prepare("SELECT a.id, a.client_id, a.balance, a.created_at, a.updated_at, c.id, c.name, c.email, c.created_at FROM accounts a INNER JOIN clients c ON a.client_id = c.id WHERE a.id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&account.ID, &account.Client.ID, &account.Balance, &account.CreatedAt, &account.UpdatedAt,
		&client.ID, &client.Name, &client.Email, &client.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}

	return &account, nil
}

func (a *AccountDB) Save(account *entity.Account) error {
	fmt.Println("Saving account:", account.ID)
	// Check if account already exists
	existingAccount, _ := a.FindByID(account.ID)

	var stmt *sql.Stmt
	var err error

	if existingAccount != nil {
		// Account exists, update it
		stmt, err = a.DB.Prepare("UPDATE accounts SET client_id = ?, balance = ?, updated_at = ? WHERE id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(account.Client.ID, account.Balance, time.Now(), account.ID)
		defer stmt.Close()
	} else {
		// Account doesn't exist, proceed with insertion
		stmt, err = a.DB.Prepare("INSERT INTO accounts (id, client_id, balance, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(account.ID, account.Client.ID, account.Balance, account.CreatedAt, time.Now())
		defer stmt.Close()
	}

	if err != nil {
		return err
	}
	return nil
}
