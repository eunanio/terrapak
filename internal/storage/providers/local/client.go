package local

import (
	"fmt"
	"os"
	"path/filepath"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"

	"github.com/gin-gonic/gin"
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

func NewProvider() *LocalProvider {
	return &LocalProvider{}
}

func (p *LocalProvider) Type() string {
	return "mnt"
}

func (p *LocalProvider) Download(mid mid.MID) (url string, err error) {
	gc := config.GetDefault()
	path := filepath.Join(gc.StorageSource.Path,WORKING_PATH, mid.Path())
	fmt.Println(path)
	return "", nil
}

func (p *LocalProvider) Upload(mid mid.MID, data []byte) error {
	gc := config.GetDefault()
	path := filepath.Join(gc.StorageSource.Path,WORKING_PATH, mid.Path())

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(mid.Path(), data, 0644); err != nil {
		return err
	}

	return nil
}

func (p *LocalProvider) Delete(mid mid.MID) error {
	return nil
}

func ServeLocalFile(c *gin.Context){
	
}