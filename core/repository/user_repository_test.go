package repository_test

import (
	"AbundancePath/core/model"
	"testing"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct {
	savedUser model.User
	foundUser model.User
}

func (m *mockUserRepository) Save(user model.User) (model.User, error) {
	m.savedUser = user
	return user, nil
}

func (m *mockUserRepository) FindByID(id int) (model.User, error) {
	return m.foundUser, nil
}

func TestSaveUser(t *testing.T) {
	repo := &mockUserRepository{}
	user := model.User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	savedUser, err := repo.Save(user)
	assert.NoError(t, err)
	assert.Equal(t, user, savedUser)
	assert.Equal(t, user, repo.savedUser)
}

func TestGetUserByID(t *testing.T) {
	repo := &mockUserRepository{}
	expectedUser := model.User{
		ID:    123,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	repo.foundUser = expectedUser

	user, err := repo.FindByID(expectedUser.ID)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}