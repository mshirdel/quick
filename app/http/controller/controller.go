package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mshirdel/quick/app"
)

type Controller struct {
	app *app.Application
}

func NewController(a *app.Application) *Controller {
	return &Controller{
		app: a,
	}
}

func (c *Controller) Routes() *echo.Echo {
	router := c.initEcho()
	// init general middleware

	return router
}

func (c *Controller) initEcho() *echo.Echo {
	r := echo.New()
	r.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "test is ok")
	})

	return r
}
