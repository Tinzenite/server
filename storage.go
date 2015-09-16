package main

/*
TODO abstract storage away so that we can write to hadoop AND disk.
Implements encrypted/Storage interface.
*/
type storage struct {
}

func createStorage() *storage {
	return &storage{}
}

func (s *storage) Store(key string, data []byte) error {
	return nil
}

func (s *storage) Retrieve(key string) ([]byte, error) {
	return nil, nil
}
