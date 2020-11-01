package repository

import (
	"sample/domain/model"
)

// ReviewRepository review repositoryのinterface
type ReviewRepository interface {
	Create(review *model.Review) (*model.Review, error)
	FindByID(id int) (*model.Review, error)
}
