package main

import (
	"log"

	"github.com/colinmarc/hdfs" // NOTE: requires the write support branch for now!
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
	fw, err := h.client.Create(key)
	if err != nil {
		// TODO test and FIXME
		// probably that is already existed... :(
		log.Println("DEBUG: probably that file already exists:", err)
	}
	// defer close for all cases
	defer fw.Close()
	// actually write data
	_, err = fw.Write(data)
	return err
}

func (h *hdfsStorage) Retrieve(key string) ([]byte, error) {
	return h.client.ReadFile(key)
}

func (h *hdfsStorage) Remove(key string) error {
	return h.client.Remove(key)
}
