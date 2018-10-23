package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/data/local"
	"net/http"
	"github.com/gericass/toilet-api/util"
	"github.com/gericass/toilet-api/handler/request"
	"database/sql"
	"github.com/gericass/toilet-api/data/remote"
)

func insertToilet(db *sql.DB, googleId string) error {
	t, err := remote.GetPlaceDetail(googleId)
	if err != nil {
		return err
	}
	toilet := ConvertPlaceDetailToToilet(t)
	err = toilet.Insert(db)
	if err != nil {
		return err
	}
	return nil

}

func GetToiletReview(c echo.Context) error {
	cc := c.(*CustomContext)
	googleId := c.Param("toiletId")
	toilet := &local.Toilet{GoogleId: googleId}
	toiletId, err := toilet.GetToiletId(cc.DB)
	if err != nil {
		return err
	}
	review := &local.Review{ToiletId: toiletId}
	reviews, err := review.FindReviewsByToiletId(cc.DB)
	if err != nil {
		return err
	}
	resp := ConvertReviews(reviews)

	return c.JSON(http.StatusOK, resp)
}

func PostReview(c echo.Context) error {
	cc := c.(*CustomContext)
	requestReview := &request.Review{}
	err := c.Bind(requestReview)
	if err != nil {
		return err
	}
	token, err := util.GetToken(c)
	if err != nil {
		return err
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(cc.DB)
	if err != nil {
		return err
	}
	toilet := &local.Toilet{GoogleId: requestReview.GoogleId}
	toiletExists, err := toilet.Exists(cc.DB)
	if err != nil {
		return err
	}
	if !toiletExists {
		err := insertToilet(cc.DB, requestReview.GoogleId)
		if err != nil {
			return err
		}
	}
	toiletId, err := toilet.GetToiletId(cc.DB)
	if err != nil {
		return err
	}
	review := &local.Review{
		UserId:    userId,
		ToiletId:  toiletId,
		Valuation: requestReview.Valuation,
		Message:   requestReview.Message,
	}
	exists, err := review.Exists(cc.DB)
	if err != nil {
		return err
	}
	if exists {
		return c.String(http.StatusCreated, "already exists")
	}
	return c.NoContent(http.StatusCreated)
}

func DeleteReview(c echo.Context) error {
	cc := c.(*CustomContext)
	googleId := c.Param("toiletId")
	token, err := util.GetToken(c)
	if err != nil {
		return err
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(cc.DB)
	if err != nil {
		return err
	}
	toilet := &local.Toilet{GoogleId: googleId}
	toiletId, err := toilet.GetToiletId(cc.DB)
	if err != nil {
		return err
	}
	review := &local.Review{ToiletId: toiletId, UserId: userId}
	err = review.DeleteReview(cc.DB)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
