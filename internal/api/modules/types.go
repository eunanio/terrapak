package modules

import "mime/multipart"

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
	file *multipart.FileHeader `form:"file"`
}