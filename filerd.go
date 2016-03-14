package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type versionType struct {
	Major string
	Minor string
	Patch string
}

var v = versionType{Major: "0", Minor: "0", Patch: "0"}

func version() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, v)
	}
}

func main() {
	e := echo.New()
	e.SetDebug(true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/", version())
	e.Post("/upload", uploadFile())
	e.Run(standard.New(":1234"))
}
