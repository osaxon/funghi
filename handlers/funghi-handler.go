package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osaxon/funghi/db"
	"github.com/osaxon/funghi/types"
)

func GetFunghi(c echo.Context) error {
	id := c.Param("id")
	d, err := db.New()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var funghi types.Funghi

	if err := d.First(&funghi, "id = ?", id).Error; err != nil {
		return c.String(http.StatusNotFound, "Funghi not found")
	}

	return c.JSON(http.StatusOK, funghi)
}

func GetFunghiList(c echo.Context) error {
	d, err := db.New()

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var funghi []types.Funghi

	if err := d.Find(&funghi).Error; err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, funghi)
}
