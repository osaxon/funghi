package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osaxon/funghi/db"
	"github.com/osaxon/funghi/handlers"
)

func main() {
	app := echo.New()
	app.Static("/static", "static")
	app.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})
	app.GET("/funghi/:id", handlers.GetFunghi)
	app.GET("/funghi", handlers.GetFunghiList)

	d, err := db.New()

	db.Migrate(d)

	if err != nil {
		app.Logger.Fatal(err)
	}

	app.Logger.Fatal(app.Start(":1323"))
}
