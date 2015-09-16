package main

import (
	"io/ioutil"

	"github.com/tinzenite/shared"
)

/*
diskStorage is a storage implementation that simply writes and reads data from
the given directory.
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
