package sqlite3

import "github.com/vkodev/filer/common"

//FileModel is a model of common.File for sqlite3 db
type FileModel struct {
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
	Provider         string
	Path             string

	ImageWidth            int
	ImageHeight           int
	ImageDateTimeOriginal string
	ImageFormat           string
	ImageGeoLocation      string
}

func (m *FileModel) TableName() string {
	return "files"
}

func (m *FileModel) ToFile(f *common.File) *common.File {
	if f == nil {
		f = &common.File{}
	}

	f.UUID = m.UUID
	f.OriginalFilename = m.OriginalFilename
	f.OriginalFileURL = m.OriginalFileURL
	f.MimeType = m.MimeType
	f.DateTimeUploaded = m.DateTimeUploaded
	f.DateTimeStored = m.DateTimeStored
	f.DateTimeDeleted = m.DateTimeDeleted
	f.IsImage = m.IsImage
	f.IsReady = m.IsReady
	f.Size = m.Size
	f.URL = m.URL

	if m.IsImage {
		f.ImageInfo = &common.ImageInfo{}
		f.ImageInfo.Width = m.ImageWidth
		f.ImageInfo.Height = m.ImageHeight
		f.ImageInfo.DateTimeOriginal = m.ImageDateTimeOriginal
		f.ImageInfo.Format = m.ImageFormat
		f.ImageInfo.GeoLocation = m.ImageGeoLocation
	}
	return f
}

func (m *FileModel) FromFile(f *common.File) {
	m.UUID = f.UUID
	m.OriginalFilename = f.OriginalFilename
	m.OriginalFileURL = f.OriginalFileURL
	m.MimeType = f.MimeType
	m.DateTimeUploaded = f.DateTimeUploaded
	m.DateTimeStored = f.DateTimeStored
	m.DateTimeDeleted = f.DateTimeDeleted
	m.IsImage = f.IsImage
	m.IsReady = f.IsReady
	m.Size = f.Size
	m.URL = f.URL

	if m.IsImage {
		f.ImageInfo = &common.ImageInfo{}
		m.ImageWidth = f.ImageInfo.Width
		m.ImageHeight = f.ImageInfo.Height
		m.ImageDateTimeOriginal = f.ImageInfo.DateTimeOriginal
		m.ImageFormat = f.ImageInfo.Format
		m.ImageGeoLocation = f.ImageInfo.GeoLocation
	}
}

func FromFile(f *common.File) *FileModel {
	m := &FileModel{}
	m.FromFile(f)
	return m
}
