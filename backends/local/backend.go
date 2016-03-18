package local

import (
	"github.com/vkodev/filer/common"
	"io"
	"os"
	"path/filepath"
	"strings"
)

//LocalStorage is an implementation of the local storage
type LocalStorage struct {
	Root string
	Name string
}

//MakeLocalStorage makes a new LocalStorage instance
func MakeLocalStorage(root string) *LocalStorage {
	if root == "" {
		root = "./storage"
	}
	return &LocalStorage{Root: root, Name: "local"}
}

// Put tries to put the file to the filesystem
func (s *LocalStorage) Write(f io.ReadCloser, file *common.File) error {
	//file.provider = s.Name
	path := makePath(s.Root, file.UUID, file.Ext())
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	//file.path = path
	// Copy
	if _, err = io.Copy(dst, f); err != nil {
		return err
	}
	return nil
}

//Read returns a file reader interface
func (s *LocalStorage) Read(file *common.File) (io.ReadCloser, error) {
	return nil, nil
}

//Remove tries to remove the file from the fs
func (s *LocalStorage) Remove(file *common.File) error {
	return nil
}

//URL returns an url to the file
func (s *LocalStorage) URL(file *common.File) (string, error) {
	return "", nil
}

//makePath generates a path to a file prefixed with a root path
//If you don't want add the root path to the path, just pass an empty string
func makePath(root string, uuid string, ext string) (path string) {
	parts := strings.SplitN(uuid, "-", 3)
	if root != "" {
		root, _ = filepath.Abs(root)
	}
	path = filepath.Join(root, filepath.Join(parts...), "original") + ext
	return
}
