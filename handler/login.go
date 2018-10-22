package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/gericass/toilet-api/data/local"
	"github.com/gericass/toilet-api/handler/request"
)

func LoginHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	i := new(request.User)
	err := c.Bind(i)
	if err != nil {
		return err
	}
	user := local.User{UID: i.UID, Name: i.Name, IconPath: i.Icon}
	exists, err := user.Exists(cc.DB)
	if err != nil {
		return err
	}
	if exists {
		return c.String(http.StatusOK, "OK")
	}
	err = user.Insert(cc.DB)
	if err != nil {
		return err
	}
	return c.String(http.StatusCreated, "OK")
}
