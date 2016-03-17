package sqlite3

import (
	"database/sql"
	"github.com/vkodev/filer/common"
)

//SqliteMetadata implementation of MetadataStorage interface
type SqliteMetadata struct {
	DB *sql.DB
}

//MakeSqliteMetadata returns instance of SqliteMetadata
func MakeSqliteMetadata(db *sql.DB) common.MetadataBackend {
	return &SqliteMetadata{DB: db}
}

//Insert new file to metadata storage
func (s *SqliteMetadata) Insert(*common.File) error {
	return nil
}

//Update file in metadata storage
func (s *SqliteMetadata) Update(*common.File) error {
	return nil
}

//Delete file from metadata storage
func (s *SqliteMetadata) Delete(*common.File) error {
	return nil
}

//Query list of files
func (s *SqliteMetadata) Query(interface{}) ([]*common.File, error) {
	return nil, nil
}
