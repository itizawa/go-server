package usecase

import (
	"sample/domain/model"
	"sample/domain/repository"
)

// ReviewUsecase review usecaseのinterface
type ReviewUsecase interface {
	Create(stars, sweetID int) (*model.Review, error)
	FindByID(id int) (*model.Review, error)
}

type reviewUsecase struct {
	reviewRepo repository.ReviewRepository
}

// NewReviewUsecase review usecaseのコンストラクタ
func NewReviewUsecase(reviewRepo repository.ReviewRepository) ReviewUsecase {
	return &reviewUsecase{reviewRepo: reviewRepo}
}

// Create reviewを保存するときのユースケース
func (ru *reviewUsecase) Create(stars, sweetID int) (*model.Review, error) {
	review, err := model.NewReview(stars, sweetID)
	if err != nil {
		return nil, err
	}

	createdReview, err := ru.reviewRepo.Create(review)
	if err != nil {
		return nil, err
	}

	return createdReview, nil
}

// FindByID reviewをIDで取得するときのユースケース
func (ru *reviewUsecase) FindByID(id int) (*model.Review, error) {
	foundReview, err := ru.reviewRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return foundReview, nil
}
