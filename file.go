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
	provider	string
	path		string
}

// ImageInfo information
type ImageInfo struct {
	Width            int
	Height           int
	DateTimeOriginal string
	Format           string
	GeoLocation      string
}

//FileRepository provides api to work with files
type FileRepository struct {
	metadata MetadataBackend
	storage FileBackend
}

//MakeNewFileRepository returns the new instance of FileRepository
func MakeNewFileRepository(storage FileBackend, metadata MetadataBackend) *FileRepository {
	return &FileRepository{metadata:metadata, storage:storage}
}
//Upload tries to upload the file and saves it to the backend
func (f *FileRepository) Upload(source io.ReadCloser, filename string) (*File, error) {
	model := MakeNewFile(filename)
	if err:=f.storage.Write(source, model); err!=nil {
		return model, err
	}

	if err := f.metadata.Insert(model); err!=nil {
		return model, err
	}

	return model, nil
}

// FileBackend is a common interface for files storage
type FileBackend interface {
	//Write tries to write the file to the backend
	Write(io.ReadCloser, *File) error
	//Read tries to get a file reader instance
	Read(*File) (io.ReadCloser, error)
	//Remove tries to remove the file from the backend
	Remove(*File) error
	//URL tries to get an url to the file
	URL(*File) (string, error)
}
// MetadataBackend is a common interface for files metadata storage
type MetadataBackend interface {
	//Insert tries to insert a new file metadata
	Insert(*File) error
	//Update tries to update a file metadata
	Update(*File) error
	//Delete tries to delete a file metadata
	Delete(*File) error
	//Query tries to get list of files metadata
	Query(interface{}) ([]*File, error)
}

// MakeNewFile make new File
func MakeNewFile(filename string) *File {
	return &File{UUID: uuid.NewV4().String(), OriginalFilename: filename, ImageInfo: &ImageInfo{}}
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

func uploadFile(fileRepository *FileRepository) echo.HandlerFunc {
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

		if err!=nil {
			return err
		}

		return c.JSON(200, f)
	}
}
