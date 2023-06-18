package main

import (
    "AbundancePath/app"
    "AbundancePath/core/model"
    "AbundancePath/core/repository"
    "fmt"
)

type userRepository struct{}

func (r *userRepository) Save(user model.User) (model.User, error) {
    return user, nil
}

func (r *userRepository) FindByID(id int) (model.User, error) {
    return model.User{ID: id, Name: "John", Email: "john@example.com"}, nil
}

type userService struct {
    repo repository.UserRepository
}

func (s *userService) CreateUser(user model.User) (model.User, error) {
    return s.repo.Save(user)
}

func (s *userService) GetUser(id int) (model.User, error) {
    return s.repo.FindByID(id)
}

func main() {
    repo := &userRepository{}
    userService := &userService{repo: repo}
    app := app.New(userService)
    err := app.Run()
    if err != nil {
        fmt.Printf("Failed to run app: %v", err)
    }
}
