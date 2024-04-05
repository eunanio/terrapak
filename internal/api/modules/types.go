package modules

import (
	"mime/multipart"
)

// List Versions DTO
type ModuleDTO struct {
	Module []ModuleVersionsDTO `json:"modules"`
}

type ModuleVersionsDTO struct {
	Versions []VersionDTO `json:"versions"`
}

type VersionDTO struct {
	Version string `json:"version"`
}

type UploadRequest struct {
	Readme string `json:"readme" form:"readme"`
	Hash   string `json:"hash" form:"hash"`
	File *multipart.FileHeader `form:"file"`
}

type UploadOptions struct {
	Readme string
	Hash   string 
	File []byte
}

type ResponseStatus struct {
	Code int `json:"status"`
	Message any `json:"message"`
}