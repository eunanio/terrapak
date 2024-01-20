package storagesource

import (
	"fmt"
	"slices"
	"strings"
)

type StorageSource struct {
	Protocol string
	Path     string
}

var (
	supportedProtocols = []string{"mnt", "s3"}
)

func NewStorageSoruce(soruce string) (StorageSource, error) {
	s := StorageSource{}
	s.Protocol = strings.Split(soruce, "://")[0]
	s.Path = strings.Split(soruce, "://")[1]
	err := validateSource(s)
	return s, err
}

func validateSource(s StorageSource) error {
	if s.Protocol == "" {
		return fmt.Errorf("[storage source error]: protocol is empty")
	}

	if s.Path == "" {
		return fmt.Errorf("[storage source error]: path is empty")
	}

	if !slices.Contains(supportedProtocols, s.Protocol) {
		return fmt.Errorf("[storage source error]: protocol %s is not supported", s.Protocol)
	}

	return nil
}

