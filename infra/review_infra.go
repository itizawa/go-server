package infra

import (
	"sample/domain/model"
	"sample/domain/repository"

	"github.com/jinzhu/gorm"
)

// ReviewRepository review repositoryの構造体
type ReviewRepository struct {
	Conn *gorm.DB
}

// NewReviewRepository review repositoryのコンストラクタ
func NewReviewRepository(conn *gorm.DB) repository.ReviewRepository {
	return &ReviewRepository{Conn: conn}
}

// Create reviewの保存
func (rr *ReviewRepository) Create(review *model.Review) (*model.Review, error) {
	if err := rr.Conn.Create(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}

// FindByID reviewをIDで取得
func (rr *ReviewRepository) FindByID(id int) (*model.Review, error) {
	review := &model.Review{ID: id}

	if err := rr.Conn.First(&review).Error; err != nil {
		return nil, err
	}

	return review, nil
}
