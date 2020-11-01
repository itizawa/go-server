package repository

import (
	"sample/domain/model"
)

// ReviewRepository review repositoryのinterface
type ReviewRepository interface {
	Create(task *model.Review) (*model.Task, error)
	FindByID(id int) (*model.Review, error)
	Update(task *model.Task) (*model.Review, error)
	Delete(task *model.Review) error
}
