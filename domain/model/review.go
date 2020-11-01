package model

import (
	"errors"
)

// Review の構造体
type Review struct {
	ID      int
	Stars   string
	SweetID string
}

// NewReview のコンストラクタ
func NewReview(stars, sweetID string) (*Review, error) {
	if stars == "" {
		return nil, errors.New("スターを入力してください")
	}

	if sweetID == "" {
		return nil, errors.New("対象の Sweet Id を入力してください")
	}

	review := &Review{
		Stars:   stars,
		SweetID: sweetID,
	}

	return review, nil
}
