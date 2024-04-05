package modules

import (
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"terrapak/internal/db/entity"
	services "terrapak/internal/db/services"
	"terrapak/internal/storage"
	"time"

	"github.com/google/uuid"
)

func Upload(m mid.MID, req UploadOptions) (res *ResponseStatus){
	gc := config.GetDefault()
	module := entity.Module{}
	ms := services.ModulesService{}
	
	storageCleint := storage.NewClient(gc.StorageSource.Protocol)
	

	moduleExsits := ms.Find(m)
	

	if req.Readme != "" {
		module.Readme = req.Readme
		if(moduleExsits.Readme != req.Readme){
			moduleExsits.Readme = req.Readme
		}
	}
	if(moduleExsits.SHAChecksum != req.Hash){
		moduleExsits.SHAChecksum = req.Hash
	}

	if moduleExsits.ID == uuid.Nil {
		module.Name = m.Name
		module.Provider = m.Provider
		module.Namespace = m.Namespace
		module.Version = m.Version
		module.SHAChecksum = req.Hash

		ms.Create(module)
	} else {
		ms.Update(moduleExsits)
	}
	

	err := storageCleint.Upload(m,req.File); if err != nil {
		res = &ResponseStatus{Code: 400, Message: "Error uploading module"}
		return res
	}
	
	res = &ResponseStatus{Code: 201, Message: "Module uploaded"}
	return res
}

func Download(m mid.MID) (res *ResponseStatus) {
	gc := config.GetDefault()
	storageClient := storage.NewClient(gc.StorageSource.Protocol)
	ms := services.ModulesService{}

	url, err := storageClient.Download(m); if err != nil {
		res = &ResponseStatus{Code: 400, Message: "Error downloading module"}
		return res
	}
	ms.IncrementDownload(m)
	res = &ResponseStatus{Code: 204, Message: url}
	return res
}

func Read(m mid.MID) *entity.Module {
	ms := services.ModulesService{}

	module := ms.Find(m); if module.ID == uuid.Nil {
		return nil
	}

	return module
}

func Versions(m mid.MID) *ResponseStatus {
	moduleDTO := ModuleDTO{}
	moduleVersionsDTO := []ModuleVersionsDTO{}
	versionDTO := []VersionDTO{}
	ms := services.ModulesService{}
	res := ResponseStatus{}
	
	list := ms.FindAll(m)
	if len(list) == 0 {
		res = ResponseStatus{Code: 404, Message: "Module not found"}
		return &res
	}

	for _, module := range list {
		versionDTO = append(versionDTO, VersionDTO{Version: module.Version})
	}
	moduleVersionsDTO = append(moduleVersionsDTO, ModuleVersionsDTO{Versions: versionDTO})
	moduleDTO.Module = moduleVersionsDTO

	res = ResponseStatus{Code: 200, Message: moduleDTO}
	return &res
}

func RemoveDraft(m mid.MID) (res *ResponseStatus) {
	gc := config.GetDefault()
	storageClient := storage.NewClient(gc.StorageSource.Protocol)
	ms := services.ModulesService{}

	module := ms.Find(m); if module.ID == uuid.Nil {
		res = &ResponseStatus{Code: 404, Message: "Module not found"} 
		return res
	}

	if module.ID != uuid.Nil {
		if module.PublishedAt.Year() < 2000 {
			ms.Remove(m)
			storageClient.Delete(m)
			res = &ResponseStatus{Code: 200, Message: "Module removed"}
			return res
		}
	}

	return nil
}

func PublishDraft(m mid.MID) (res *ResponseStatus) {
	ms := services.ModulesService{}
	module := ms.Find(m); if module.ID == uuid.Nil {

		res = &ResponseStatus{Code: 404, Message: "Module not found"}
		return res
	}

	if module.ID != uuid.Nil && module.PublishedAt.IsZero() {
		module.PublishedAt = time.Now()
		ms.Update(module)
		res = &ResponseStatus{Code: 200, Message: "Module published"}
	}

	return nil
}