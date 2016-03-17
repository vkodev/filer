package main

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"log"
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
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	metadata := MakeSqliteMetadata(db)
	storage := MakeLocalStorage("")
	repository := MakeNewFileRepository(storage, metadata)

	e := echo.New()
	e.SetDebug(true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/", version())
	e.Post("/upload", uploadFile(repository))
	e.Run(standard.New(":1234"))
}
