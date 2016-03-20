package common

import (
	"io"
	"time"
)

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

type ApiTokenRepository interface {
	FindByToken(token string) (*ApiToken, error)

	FindOneByUser(userId int) (*ApiToken, error)

	Create(apiToken *ApiToken) error

	UpdateExpiry(token string, expiry time.Time) error

	Delete(token string) error
}

type ApiUserRepository interface {
	FindById(id int) (*ApiUser, error)

	FindByLogin(login string) (*ApiUser, error)

	FindAll() ([]*ApiUser, error)

	Create(apiUser *ApiUser) error

	Update(apiUser *ApiUser) error

	Delete(id int) error
}
