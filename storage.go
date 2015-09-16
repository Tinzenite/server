package main

import "github.com/tinzenite/shared"

/*
TODO abstract storage away so that we can write to hadoop AND disk.
Implements encrypted/Storage interface.
*/
type storage struct {
	useHadoop bool
}

/*
createStorage builds a storage interface.
TODO useHadoop is currently not used!
*/
func createStorage(useHadoop bool) *storage {
	return &storage{useHadoop: false}
}

func (s *storage) Store(key string, data []byte) error {
	if s.useHadoop {
		return shared.ErrUnsupported
	}
	return nil
}

func (s *storage) Retrieve(key string) ([]byte, error) {
	if s.useHadoop {
		return nil, shared.ErrUnsupported
	}
	return nil, nil
}
