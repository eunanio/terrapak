package router

import (
	"bytes"
	"io"
	"terrapak/internal/api/modules"
	"terrapak/internal/config/mid"

	"github.com/gin-gonic/gin"
)

type Endpoint struct{}

func (e *Endpoint) Read(c *gin.Context) {
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
	}

	module := modules.Read(m)
	if module == nil {
		c.JSON(404, gin.H{
			"message": "Module not found",
		})
		return
	}

	c.JSON(200, module)
}

func (e *Endpoint) Download(c *gin.Context) {
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
	}

	status := modules.Download(m)
	c.Status(status.Code)
	c.Header("X-Terraform-Get", status.Message.(string))
}

func (e *Endpoint) Version(c *gin.Context) {
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
	}

	status := modules.Versions(m)
	c.JSON(status.Code, gin.H{
		"message": status.Message,
	})
}

func (e *Endpoint) Publish(c *gin.Context) {
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
	}

	status := modules.PublishDraft(m)
	c.JSON(status.Code, gin.H{
		"message": status.Message,
	})
}

func (e *Endpoint) Remove(c *gin.Context) {
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
	}

	status := modules.RemoveDraft(m)
	c.JSON(status.Code, gin.H{
		"message": status.Message,
	})
}

func (e *Endpoint) Upload(c *gin.Context) {
	req := modules.UploadRequest{}
	uploadOpts := modules.UploadOptions{Readme: req.Readme, Hash: req.Hash}
	m, err := mid.Parse(c); if err != nil {
		c.JSON(400, err)
	}
	buffer := bytes.NewBuffer(nil)
	err = c.Bind(&req); if err != nil {
		c.JSON(400, err)
		return
	}

	src, err := req.File.Open(); if err != nil {
		c.JSON(400, err)	
	}
	defer src.Close()
	io.Copy(buffer, src)
	uploadOpts.File = buffer.Bytes()
	
	status := modules.Upload(m, uploadOpts)
	c.JSON(status.Code, gin.H{
		"message": status.Message,
	})
}