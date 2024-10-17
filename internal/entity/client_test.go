package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john.doe@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john.doe@example.com", client.Email)
}

func TestNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdate(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")

	err := client.Update("John Doe Updated", "john.doe.updated@example.com")

	assert.Nil(t, err)
	assert.Equal(t, "John Doe Updated", client.Name)
	assert.Equal(t, "john.doe.updated@example.com", client.Email)
}

func TestUpdateWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")

	err := client.Update("", "john.doe.updated@example.com")

	assert.NotNil(t, err)
	assert.Equal(t, "name and email are required", err.Error())
}

func TestAddAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john.doe@example.com")
	account, _ := NewAccount(client)

	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
