package controller

import (
	"net/http"

	"../models"

	"github.com/labstack/echo"
)

func GetKategori(c echo.Context) error {

	result := models.GetKat()
	return c.JSON(http.StatusOK, result)
}
