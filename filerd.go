package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vkodev/filer/backends/local"
	"github.com/vkodev/filer/backends/sqlite3"
	"github.com/vkodev/filer/common"
	"github.com/vkodev/filer/rest-api"
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
	db, err := gorm.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	metadata := sqlite3.MakeSqliteMetadata(db)
	storage := local.MakeLocalStorage("")
	repository := common.MakeNewFileRepository(storage, metadata)

	// TODO: we need use better abstraction with choice from config, like if ... InitGormRepositories
	tokenRepository := common.MakeApiTokenRepository(db)
	userRepository := common.MakeApiUserGormRepository(db)

	e := echo.New()
	e.SetDebug(true)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/", version())
	e.Post("/upload", common.UploadFileHandler(repository))
	e.Post("/api/v1/auth", rest_api.HandleAuth(tokenRepository, userRepository))
	e.Run(standard.New(":1234"))
}
