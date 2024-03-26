package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountJSONSerialization(t *testing.T) {
	account := NewAccount("dummy", "dummy@example.com", "password123")

	data, err := json.Marshal(account)
	assert.NoError(t, err)

	var newAccount Account
	err = json.Unmarshal(data, &newAccount)
	assert.NoError(t, err)

	assert.Equal(t, account.Name, newAccount.Name)
	assert.Equal(t, account.Email, newAccount.Email)
	assert.Equal(t, account.Password, newAccount.Password)
	assert.Equal(t, account.AvatarURL, newAccount.AvatarURL)
	assert.Equal(t, account.Role, newAccount.Role)
}
