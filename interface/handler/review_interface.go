package handler

import (
	"net/http"
	"strconv"

	"sample/usecase"

	"github.com/labstack/echo"
)

// ReviewHandler review handlerのinterface
type ReviewHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
}

// ReviewHandler review handlerの構造体
type ReviewHandler struct {
	reviewUsecase usecase.ReviewUsecase
}

// NewReviewHandler review handlerのコンストラクタ
func NewReviewHandler(reviewUsecase usecase.ReviewUsecase) ReviewHandler {
	return &reviewHandler{reviewUsecase: reviewUsecase}
}

type requestReview struct {
	Stars   int `json:"stars"`
	SweetID int `json:"sweetID"`
}

type responseReview struct {
	ID      int `json:"id"`
	Stars   int `json:"stars"`
	SweetID int `json:"sweetID"`
}

// Post reviewを保存するときのハンドラー
func (rh *reviewHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdReview, err := rh.reviewUsecase.Create(req.stars, req.sweetID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseReview{
			ID:      createdReview.ID,
			Stars:   createdReview.stars,
			SweetID: createdReview.sweetID,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

// Get reviewを取得するときのハンドラー
func (rh *reviewHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundReview, err := rh.reviewUsecase.FindByID(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseReview{
			ID:      foundReview.ID,
			Stars:   foundReview.Stars,
			SweetID: foundReview.SweetID,
		}

		return c.JSON(http.StatusOK, res)
	}
}
