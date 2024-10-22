package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com.br/devfullcycle/fc-ms-wallet/internal/database"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_account"
	"github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	clientDB := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)

	// Create tables
	createTables(db)

	// Create a client
	clientUseCase := create_client.NewCreateClientUseCase(clientDB)
	client, err := clientUseCase.Execute(create_client.CreateClientInputDTO{
		Name:  "John Doe",
		Email: "john@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Client created: %+v\n", client)

	// Create an account
	accountUseCase := create_account.NewCreateAccountUseCase(accountDB, clientDB)
	account, err := accountUseCase.Execute(create_account.CreateAccountInputDTO{
		ClientID: client.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account created: %+v\n", account)

	// Create another account for transaction
	account2, err := accountUseCase.Execute(create_account.CreateAccountInputDTO{
		ClientID: client.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account created: %+v\n", account2)

	// Deposit some money to the first account
	fmt.Println("Depositing money to the first account", account.ID)
	acc, err := accountDB.FindByID(account.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Account found: %+v\n", acc)
	err = acc.Deposit(1000)
	if err != nil {
		log.Fatal(err)
	}
	err = accountDB.Save(acc)
	if err != nil {
		log.Fatal(err)
	}

	// Check final balances
	acc1, _ := accountDB.FindByID(account.ID)
	acc2, _ := accountDB.FindByID(account2.ID)
	fmt.Printf("Account 1 balance: %f\n", acc1.Balance)
	fmt.Printf("Account 2 balance: %f\n", acc2.Balance)

	dropTables(db)
}

func createTables(db *sql.DB) {
	db.Exec(`CREATE TABLE IF NOT EXISTS clients (
		id VARCHAR(255) NOT NULL,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME,
		PRIMARY KEY (id)
	)`)

	db.Exec(`CREATE TABLE IF NOT EXISTS accounts (
		id VARCHAR(255) NOT NULL,
		client_id VARCHAR(255) NOT NULL,
		balance FLOAT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME,
		PRIMARY KEY (id),
		FOREIGN KEY (client_id) REFERENCES clients(id)
	)`)
}

func dropTables(db *sql.DB) {
	db.Exec(`DROP TABLE IF EXISTS clients`)
	db.Exec(`DROP TABLE IF EXISTS accounts`)
}
