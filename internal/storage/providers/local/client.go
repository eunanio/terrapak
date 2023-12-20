package local

import (
	"os"
	"path/filepath"
	"terrapak/internal/config/mid"
)

const (
	WORKING_PATH = "/terrapak/modules/"
)

type LocalProvider struct {
	config LocalConfig
}

type LocalConfig struct {
	// Path to the local storage directory
	Path string `yaml:"path"`
}

func NewProvider() {

}

func Type() string {
	return "local"
}

func Download(mid mid.MID) (url string, err error) {
	return "", nil
}

func Upload(mid mid.MID, data []byte) error {
	path := filepath.Join(WORKING_PATH, mid.Path())

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(mid.Path(), data, 0644); err != nil {
		return err
	}

	return nil
}