package modules

import (
	"bytes"
	"fmt"
	"io"
	"terrapak/internal/api/metadata"
	"terrapak/internal/config"
	"terrapak/internal/config/mid"
	"terrapak/internal/db/entity"
	services "terrapak/internal/db/services"
	"terrapak/internal/storage"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Upload(c *gin.Context) {
	gc := config.GetDefault()
	module := entity.Module{}
	ms := services.ModulesService{}
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
		return
	}
	storageCleint := storage.NewClient(gc.StorageSource.Protocol)
	buffer := bytes.NewBuffer(nil)
	
	moduleExsits := ms.Find(m)
	if c.PostForm("readme") != "" {
		module.Readme = c.PostForm("readme")
		fmt.Println(module.Readme)
	}

	if moduleExsits.ID == uuid.Nil {
		module.Name = m.Name
		module.Provider = m.Provider
		module.Namespace = m.Namespace
		module.Version = m.Version

		ms.Create(module)
	}
	
	file, err := c.FormFile("file"); if err != nil {
		panic(err)
	}

	src,_ := file.Open()
	defer src.Close()
	io.Copy(buffer, src)

	storageCleint.Upload(m,buffer.Bytes())
}

func Download(c *gin.Context) {
	gc := config.GetDefault()
	storageClient := storage.NewClient(gc.StorageSource.Protocol)
	ms := services.ModulesService{}
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
		return
	}

	url, err := storageClient.Download(m); if err != nil {
		c.JSON(400, err)
		return
	}
	ms.IncrementDownload(m)
	c.Header("X-Terraform-Get", url)
	c.Status(204)
}

func Read(c *gin.Context) {
	ms := services.ModulesService{}
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400,metadata.NewApiResponse(400, err.Error()))
		return
	}

	module := ms.Find(m); if module.ID == uuid.Nil {
		c.JSON(404,metadata.NewApiResponse(404, "module not found"))
		return
	}

	c.JSON(200,module)
}

func Versions(c *gin.Context) {
	moduleDTO := ModuleDTO{}
	moduleVersionsDTO := []ModuleVersionsDTO{}
	versionDTO := []VersionDTO{}
	ms := services.ModulesService{}
	m, err := mid.Parse(c); if err != nil {
		fmt.Println(err)
		errResposne := metadata.ApiResponse{Code: 400, Message: err.Error()}
		c.JSON(400,errResposne)
		return
	}
	
	list := ms.FindAll(m)
	fmt.Println(list)
	if len(list) == 0 {
		err := metadata.ApiResponse{Code: 404, Message: "module not found"}
		c.JSON(404,err)
		return
	}

	for _, module := range list {
		versionDTO = append(versionDTO, VersionDTO{Version: module.Version})
	}
	moduleVersionsDTO = append(moduleVersionsDTO, ModuleVersionsDTO{Versions: versionDTO})
	moduleDTO.Module = moduleVersionsDTO

	c.JSON(200,moduleDTO)
}

func RemoveDraft(c *gin.Context) {
	gc := config.GetDefault()
	storageClient := storage.NewClient(gc.StorageSource.Protocol)

	m := mid.MID{}
	ms := services.ModulesService{}
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400,metadata.NewApiResponse(400, "Error Parsing MID"))
		return
	}

	module := ms.Find(m); if module.ID == uuid.Nil {
		c.JSON(404,metadata.NewApiResponse(404, "Module not found"))
		return
	}

	if module.ID != uuid.Nil {
		if module.PublishedAt.Year() < 2000 {
			ms.Remove(m)
			storageClient.Delete(m)
			c.JSON(200,metadata.NewApiResponse(200,"Module deleted"))
		}
	}
}

func PublishDraft(c *gin.Context) {
	m := mid.MID{}
	ms := services.ModulesService{}
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400,metadata.NewApiResponse(400, err.Error()))
		return
	}
	module := ms.Find(m); if module.ID == uuid.Nil {
		c.JSON(404,metadata.NewApiResponse(404, "Module not found"))
		return
	}

	if module.ID != uuid.Nil {
		module.PublishedAt = time.Now()
		ms.Update(module)
	}
}