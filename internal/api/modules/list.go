package modules

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