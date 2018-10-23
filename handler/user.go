package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/data/local"
	"net/http"
	"github.com/gericass/toilet-api/handler/response"
)

func GetUserHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	uid := c.Param("uid")
	user := &local.User{UID: uid}
	err := user.FindUserByUID(cc.DB)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func GetUserReviewHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	uid := c.Param("uid")
	user := &local.User{UID: uid}
	userId, err := user.GetUserId(cc.DB)
	if err != nil {
		return err
	}
	review := &local.Review{UserId: userId}
	exists, err := review.ExistsByUserId(cc.DB)
	if err != nil {
		return err
	}
	if !exists {
		r := new(response.Reviews)
		r.Status = "ReviewEmpty"
		return c.JSON(http.StatusOK, r)
	}
	localReviews, err := review.FindReviewsByUserId(cc.DB)
	if err != nil {
		return err
	}
	resp := ConvertReviews(localReviews, cc.DB)
	return c.JSON(http.StatusOK, resp)
}
