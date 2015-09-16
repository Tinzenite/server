package main

import (
	"io/ioutil"

	"github.com/tinzenite/shared"
)

/*
TODO abstract storage away so that we can write to hadoop AND disk.
Implements encrypted/Storage interface.
*/
type diskStorage struct {
	RootPath string
}

/*
createDiskStorage creates a structure that writes all files to disk.
*/
func createDiskStorage(path string) *diskStorage {
	return &diskStorage{RootPath: path}
}

func (s *diskStorage) Store(key string, data []byte) error {
	return ioutil.WriteFile(s.RootPath+"/"+key, data, shared.FILEPERMISSIONMODE)
}

func (s *diskStorage) Retrieve(key string) ([]byte, error) {
	return ioutil.ReadFile(s.RootPath + "/" + key)
}
