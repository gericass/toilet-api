package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/data/remote"
	"github.com/gericass/toilet-api/handler/response"
	"net/http"
	"time"
	"database/sql"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/util"
)

func assignDescription(db *sql.DB, t *response.Toilet) error {
	toilet := &local.Toilet{GoogleId: t.GoogleId}
	exists, err := toilet.Exists(db)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	err = toilet.FindToiletByGoogleId(db)
	if err != nil {
		return err
	}
	t.Name = toilet.Name
	t.Image = toilet.ImagePath
	t.Description = toilet.Description
	t.Valuation = toilet.Valuation
	t.UpdatedAt = toilet.UpdatedAt
	return nil
}

func assignFavorite(c echo.Context, db *sql.DB, t *response.Toilet) error {
	toilet := &local.Toilet{GoogleId: t.GoogleId}
	exists, err := toilet.Exists(db)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	toilet.FindToiletByGoogleId(db)
	token, err := util.GetToken(c)
	if err != nil {
		return err
	}
	user := &local.User{UID: token.UID}
	userId, err := user.GetUserId(db)
	if err != nil {
		return err
	}
	usersToilets := &local.UsersToilets{ToiletId: toilet.ID, UserId: userId}
	favorite, err := usersToilets.Exists(db)
	if err != nil {
		return err
	}
	if favorite {
		t.Favorite = true
		return nil
	}
	t.Favorite = false
	return nil
}

func createToilets(c echo.Context, places *remote.Place, db *sql.DB) *response.Toilets {
	var ts []*response.Toilet
	for _, v := range places.Results {
		t := &response.Toilet{Name: v.Name, GoogleId: v.PlaceID, Geolocation: v.Vicinity, Image: v.Icon, UpdatedAt: time.Now()}
		assignDescription(db, t)
		assignFavorite(c, db, t)
		ts = append(ts, t)
	}
	return &response.Toilets{Toilets: ts}
}

func SearchHandler(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	cc := c.(*CustomContext)
	places, err := remote.SearchPlaces(keyword)
	if err != nil {
		return err
	}
	toilets := createToilets(c, places, cc.DB)
	return c.JSON(http.StatusOK, toilets)
}
