package sqlite3

import (
	"github.com/jinzhu/gorm"
	"github.com/vkodev/filer/common"
)

//SqliteMetadata implementation of MetadataStorage interface
type SqliteMetadata struct {
	DB *gorm.DB
}

//MakeSqliteMetadata returns instance of SqliteMetadata
func MakeSqliteMetadata(db *gorm.DB) common.MetadataBackend {
	//TODO: write the more clever migration
	db.AutoMigrate(&FileModel{})
	return &SqliteMetadata{DB: db}
}

//Insert new file to metadata storage
func (s *SqliteMetadata) Insert(f *common.File) error {
	m := FromFile(f)
	s.DB.Create(m)
	m.ToFile(f)
	return nil
}

//Update file in metadata storage
func (s *SqliteMetadata) Update(f *common.File) error {
	m := FromFile(f)
	s.DB.Update(m)
	m.ToFile(f)
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

func (s *SqliteMetadata) findByUUID(uuid string) (*FileModel, error) {
	m := &FileModel{}
	s.DB.Where("uuid = ?", uuid).First(m)
	if s.DB.RecordNotFound() {
		return nil, nil
	}
	return m, nil
}
