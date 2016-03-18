package common

import "github.com/labstack/echo"

//UploadFileHandler handles files upload
func UploadFileHandler(fileRepository *FileRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		//store:= c.Form("store")
		filename := c.Form("filename")
		file, err := req.FormFile("file")
		if err != nil {
			return err
		}

		if filename == "" {
			filename = file.Filename
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		f, err := fileRepository.Upload(src, filename)

		if err != nil {
			return err
		}

		return c.JSON(200, f)
	}
}
