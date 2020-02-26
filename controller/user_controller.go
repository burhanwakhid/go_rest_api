package controller

import (
	// "net/http"

	"../models"

	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
)

func GetAllUser(c echo.Context) error {
	result := models.GetUser()
	return c.JSON(fasthttp.StatusOK, result)
}
