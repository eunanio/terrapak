package services

import (
	"terrapak/internal/config/mid"
	"terrapak/internal/db/client"
	"terrapak/internal/db/entity"
	"time"
)

type ModulesService struct{}

func (ms *ModulesService) Create(module entity.Module) {
	client := client.Default()
	module.DownloadCount = 0
	
	module.Create(client, &module)
}

func (ms *ModulesService) Update(module *entity.Module) {
	client := client.Default()
	module.UpdatedAt = time.Now()
	module.Update(client, module)
}

func (ms *ModulesService) Find(mid mid.MID) *entity.Module {
	model := entity.Module{}
	client := client.Default()

	return model.Read(client, mid)
}

func (ms *ModulesService) FindAll(mid mid.MID) []entity.Module {
	model := entity.Module{}
	client := client.Default()

	return model.ReadAll(client, mid)
}

func (ms *ModulesService) IncrementDownload(mid mid.MID) {
	model := entity.Module{}
	client := client.Default()
	model = *model.Read(client, mid)
	model.DownloadCount = model.DownloadCount + 1
	model.Update(client, &model)
}

func (ms *ModulesService) Remove(mid mid.MID) {
	model := entity.Module{}
	client := client.Default()
	model = *model.Read(client, mid)
	model.Delete(client, &model)
}