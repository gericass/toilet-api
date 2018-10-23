package main

import (
	"github.com/labstack/echo"
	"github.com/gericass/toilet-api/handler"
	"github.com/gericass/toilet-api/data/local"
	"github.com/labstack/echo/middleware"
	"net/http"
	"github.com/gericass/toilet-api/util"
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

func validateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		forbidden := c.String(http.StatusForbidden, "Forbidden")
		_, err := util.GetToken(c)
		if err != nil {
			return forbidden
		}
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(dbMiddleware)
	e.Use(validateToken)
	e.Use(middleware.Logger())

	e.POST("/login", handler.LoginHandler)

	e.GET("/search", handler.SearchHandler)

	e.GET("/favorite", handler.GetFavoriteHandler)
	e.POST("/favorite", handler.PostFavoriteHandler)
	e.DELETE("/favorite/:toiletId", handler.DeleteFavoriteHandler)

	e.GET("/review/:toiletId", handler.GetToiletReview)
	e.POST("/review", handler.PostReview)
	e.DELETE("/review/:toiletId", handler.DeleteReview)

	e.GET("/user/:uid")

	e.Start(":8000")
}
