package common

import (
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
	provider         string
	path             string
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
	storage  FileBackend
}

//MakeNewFileRepository returns the new instance of FileRepository
func MakeNewFileRepository(storage FileBackend, metadata MetadataBackend) *FileRepository {
	return &FileRepository{metadata: metadata, storage: storage}
}

//Upload tries to upload the file and saves it to the backend
func (f *FileRepository) Upload(source io.ReadCloser, filename string) (*File, error) {
	model := MakeNewFile(filename)
	if err := f.storage.Write(source, model); err != nil {
		return model, err
	}

	if err := f.metadata.Insert(model); err != nil {
		return model, err
	}

	return model, nil
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
