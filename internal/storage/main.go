package storage

import "terrapak/internal/config/mid"

type StorageProvider interface {
	NewProvider() StorageProvider
	Type() string
	Download(mid mid.MID) (url string, err error)
	Upload(mid mid.MID, data []byte) error
}

type Storage struct {
	Provider StorageProvider
}


