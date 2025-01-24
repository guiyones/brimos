package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Guilherme", "guiyonesnogara@gmail.com", "gui070591")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Guilherme", user.Name)
	assert.Equal(t, "guiyonesnogara@gmail.com", user.Email)
}

func TestUser_ValidatePasswo(t *testing.T) {
	user, err := NewUser("Guilherme", "guiyonesnogara@gmail.com", "gui070591")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("gui070591"))
	assert.False(t, user.ValidatePassword("gui07059"))
	assert.NotEqual(t, "gui070591", user.Password)
}
