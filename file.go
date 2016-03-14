package main

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"io"
	"path/filepath"
)

// File representation
type File struct {
	UUID             string
	OriginalFilename string
	OriginalFileURL  string
	MimeType         string
	DateTimeUploaded string
	DateTimeStored   string
	DateTimeDeleted  string
	IsImage          bool
	IsReady          bool
	Size             int
	URL              string
	ImageInfo        *ImageInfo
	fileInfo         *FileInfo
}

// ImageInfo information
type ImageInfo struct {
	Width            int
	Height           int
	DateTimeOriginal string
	Format           string
	GeoLocation      string
}

// FileInfo additional info about file
type FileInfo struct {
	Provider string
	Path     string
}

// Storage common interface for storage
type Storage interface {
	Put(io.ReadCloser, *File) error
	Get(*File) (io.ReadCloser, error)
	Remove(*File) error
}

// NewFile make new File
func NewFile() *File {
	return &File{UUID: uuid.NewV4().String(), ImageInfo: &ImageInfo{}, fileInfo: &FileInfo{}}
}

//Ext return file ext
func (f *File) Ext() string {
	if f.IsImage {
		return f.ImageInfo.Format
	}

	if f.OriginalFilename != "" {
		return filepath.Ext(f.OriginalFilename)
	}
	return filepath.Ext(f.OriginalFileURL)
}

func uploadFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		storage := MakeLocalStorage("")
		//store:= c.Form("store")
		filename := c.Form("filename")
		f := NewFile()

		file, err := req.FormFile("file")
		if err != nil {
			return err
		}

		if filename != "" {
			f.OriginalFilename = filename
		} else {
			f.OriginalFilename = file.Filename
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		if err = storage.Put(src, f); err != nil {
			return err
		}
		return c.JSON(200, f)
	}
}
