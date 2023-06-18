package service

import "AbundancePath/core/model"

type UserService interface {
    CreateUser(user model.User) (model.User, error)
    GetUser(id int) (model.User, error)
}
