package entity

import (
	"fmt"
	"terrapak/internal/config/mid"
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ModelBase
	Name          string    `json:"name"`
	Provider      string    `json:"provider"`
	Namespace     string    `json:"namespace"`
	Version       string    `json:"version"`
	DownloadCount int       `json:"download_count"`
	PublishedAt   time.Time `json:"published_at"`
	Readme        string    `json:"readme"`
}

func (Module) TableName() string {
	return "modules"
}

func (m *Module) Up(client *gorm.DB) {
	fmt.Println("[SETUP] - creating modules table")
	err := client.AutoMigrate(&Module{}); if err != nil {
		panic("error migrating modules table")
	}
}

func (m *Module) Create(client *gorm.DB, module *Module) {
	client.Create(module)
}

func (m *Module) Read(client *gorm.DB, mid mid.MID) (module *Module) {
	result := client.Where("Namespace = ? AND Provider = ? AND Name = ? AND Version = ?", mid.Namespace,mid.Provider,mid.Name,mid.Version).First(&module); if result.Error != nil {
		fmt.Println(result.Error)
	}
	return module
}

func (m *Module) Update(client *gorm.DB, module *Module) {
	module.UpdatedAt = time.Now()
	client.Save(module)
}

func (m *Module) Delete(client *gorm.DB, module *Module) {
	client.Delete(&module)
}

func (m *Module) ReadAll(db *gorm.DB, mid mid.MID) []Module {
	list := []Module{}
	m.Namespace = mid.Namespace
	m.Provider = mid.Provider
	m.Name = mid.Name
	db.Find(&list,m)
	return list
}