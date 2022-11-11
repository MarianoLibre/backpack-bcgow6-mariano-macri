package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"library.com/db"
	"library.com/internal/domain"
)

func TestSave(t *testing.T) {
	user := domain.User{
		Name: "TestUser",
		Age: 18,
	}

	db.Init()
	testRepo := NewRepository(db.DataBase)
	result, err := testRepo.Save(user)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, result.Name)
	assert.Equal(t, user.Age, result.Age)
}
