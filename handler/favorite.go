package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/util"
	"net/http"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/handler/request"
	"github.com/gericass/toilet-api/data/remote"
	"github.com/gericass/toilet-api/handler/response"
)

func GetFavoriteHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	token, err := util.GetToken(c)
	if err != nil {
		return c.String(http.StatusForbidden, "Forbidden")
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(cc.DB)
	if err != nil {
		return err
	}
	usersToilets := &local.UsersToilets{UserId: userId}
	exists, err := usersToilets.ExistsByUserId(cc.DB)
	if err != nil {
		return err
	}
	if !exists {
		r := new(response.Reviews)
		r.Status = "FavoriteEmpty"
		return c.JSON(http.StatusOK, r)
	}
	usersToiletsList, err := usersToilets.FindToiletsByUserId(cc.DB)
	if err != nil {
		return err
	}
	var toilets []*local.Toilet
	for _, v := range usersToiletsList {
		t := &local.Toilet{ID: v.ToiletId}
		t.FindToiletByGoogleId(cc.DB)
		toilets = append(toilets, t)
	}
	resp := ConvertToilets(toilets)
	return c.JSON(http.StatusOK, resp)
}

func PostFavoriteHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	requestFavorite := new(request.Favorite)
	err := c.Bind(requestFavorite)
	if err != nil {
		return err
	}
	toilet := &local.Toilet{GoogleId: requestFavorite.GoogleId}
	toiletExists, err := toilet.Exists(cc.DB)
	if err != nil {
		return err
	}
	if !toiletExists {
		placeDetail, err := remote.GetPlaceDetail(toilet.GoogleId)
		if err != nil {
			return err
		}
		t := ConvertPlaceDetailToToilet(placeDetail)
		err = t.Insert(cc.DB)
		if err != nil {
			return err
		}
	}

	err = toilet.FindToiletByGoogleId(cc.DB)
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
	usersToilets := &local.UsersToilets{ToiletId: toilet.ID, UserId: userId}
	exists, err := usersToilets.Exists(cc.DB)
	if err != nil {
		return err
	}
	if exists {
		return c.String(http.StatusConflict, "already exists")
	}
	err = usersToilets.Insert(cc.DB)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusCreated)
}

func DeleteFavoriteHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	placeId := c.Param("toiletId")
	toilet := &local.Toilet{GoogleId: placeId}
	toiletId, err := toilet.GetToiletId(cc.DB)
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
	usersToilets := &local.UsersToilets{ToiletId: toiletId, UserId: userId}
	err = usersToilets.DeleteUsersToilets(cc.DB)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}
