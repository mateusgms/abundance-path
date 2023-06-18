package service_test

import (
	"AbundancePath/core/model"
	"AbundancePath/core/repository"
	"testing"
	"github.com/stretchr/testify/assert"
)

type mockUserService struct {
	repo repository.UserRepository
}

func (m *mockUserService) CreateUser(user model.User) (model.User, error) {
	return user, nil
}

func (m *mockUserService) GetUser(id int) (model.User, error) {

	mockUser := model.User{
		ID:    id,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	return mockUser, nil
}
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
func TestCreateUser(t *testing.T) {
	repo := &mockUserRepository{}
	svc := &mockUserService{repo: repo}

	user := model.User{
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	createdUser, err := svc.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)
}

func TestGetUser(t *testing.T) {
	repo := &mockUserRepository{}
	svc := &mockUserService{repo: repo}

	userID := 123

	user, err := svc.GetUser(userID)
	assert.NoError(t, err)
	assert.Equal(t, userID, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "johndoe@example.com", user.Email)
}
