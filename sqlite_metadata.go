package main

import "database/sql"

//SqliteMetadata implementation of MetadataStorage interface
type SqliteMetadata struct {
	DB *sql.DB
}

//MakeSqliteMetadata returns instance of SqliteMetadata
func MakeSqliteMetadata(db *sql.DB) MetadataBackend {
	return &SqliteMetadata{DB: db}
}

//Insert new file to metadata storage
func (s *SqliteMetadata) Insert(*File) error {
	return nil
}

//Update file in metadata storage
func (s *SqliteMetadata) Update(*File) error {
	return nil
}

//Delete file from metadata storage
func (s *SqliteMetadata) Delete(*File) error {
	return nil
}

//Query list of files
func (s *SqliteMetadata) Query(interface{}) ([]*File, error) {
	return nil, nil
}
