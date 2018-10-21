package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/handler"
	"github.com/gericass/toilet-api/data/local"
	"github.com/labstack/echo/middleware"
)

func dbMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := local.ConnectDB()
		if err != nil {
			return err
		}
		defer db.Close()
		cc := &handler.CustomContext{c, db}
		return h(cc)
	}
}

func main() {
	e := echo.New()
	e.Use(dbMiddleware)
	e.Use(middleware.Logger())

	e.POST("/login", handler.LoginHandler)

	e.Start(":8000")
}
