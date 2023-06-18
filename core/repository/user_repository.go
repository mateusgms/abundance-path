package repository

import "AbundancePath/core/model"

type UserRepository interface {
	Save(user model.User) (model.User, error)
	FindByID(id int) (model.User, error)
}