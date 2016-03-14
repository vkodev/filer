package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

//LocalStorage implementation of local storage
type LocalStorage struct {
	Root string
	Name string
}

//MakeLocalStorage make new LocalStorage instance
func MakeLocalStorage(root string) *LocalStorage {
	if root == "" {
		root = "./storage"
	}
	return &LocalStorage{Root: root, Name: "local"}
}

// Put file to fs
func (s *LocalStorage) Put(f io.ReadCloser, file *File) error {
	file.fileInfo.Provider = s.Name
	path := makePath(s.Root, file.UUID, file.Ext())
	if err := os.MkdirAll(filepath.Dir(path), 0777); err!=nil {
		return err;
	}
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	file.fileInfo.Path = path
	// Copy
	if _, err = io.Copy(dst, f); err != nil {
		return err
	}
	return nil
}

//Get file instance
func (s *LocalStorage) Get(file *File) (io.ReadCloser, error) {
	return nil, nil
}

//Remove file from fs
func (s *LocalStorage) Remove(file *File) error {
	return nil
}

func makePath(root string, uuid string, ext string) (path string) {
	parts := strings.SplitN(uuid, "-", 3)
	root, _ = filepath.Abs(root)
	path = filepath.Join(root, filepath.Join(parts...),"original") + ext
	return
}
