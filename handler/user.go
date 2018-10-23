package handler

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/data/local"
)

func GetUserHandler(c echo.Context) error {
	cc := c.(*CustomContext)
	uid := c.Param("uid")
	user := &local.User{UID: uid}
	err := user.FindUserByUID(cc.DB)
	if err != nil {
		return err
	}

	return nil
}
