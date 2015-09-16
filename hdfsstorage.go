package main

import (
	"github.com/colinmarc/hdfs"
	"github.com/tinzenite/shared"
)

/*
hdfsStorage is a storage implementation that reads and writes data via a Hadoop
file system.

Implements encrypted/Storage interface.
*/
type hdfsStorage struct {
	client *hdfs.Client
}

/*
createHDFSStorage creates a structure that writes all files to a Hadoop file system.
*/
func createHDFSStorage(url string) (*hdfsStorage, error) {
	// connect to URL
	client, err := hdfs.New(url) // FIXME: use NewForUser()
	if err != nil {
		return nil, err
	}
	return &hdfsStorage{
		client: client}, nil
}

func (h *hdfsStorage) Store(key string, data []byte) error {
	return shared.ErrUnsupported
}

func (h *hdfsStorage) Retrieve(key string) ([]byte, error) {
	return h.client.ReadFile(key)
}
