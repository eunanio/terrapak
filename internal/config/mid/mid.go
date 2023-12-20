package mid

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type MID struct {
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Namespace 	string `json:"namespace"`
	Version     string `json:"version"`
}

func NewMID(name, provider, namespace, version string) MID {
	return MID{
		Name: name,
		Provider: provider,
		Namespace: namespace,
		Version: version,
	}
}

func Parse(c *gin.Context) (mid MID, err error){
	mid.Name = c.Param("name")
	mid.Provider = c.Param("provider")
	mid.Namespace = c.Param("namespace")
	mid.Version = c.Param("version")
	err = validate(mid)
	return mid, err
}

func validate(mid MID) error {

	if mid.Name == "" {
		return fmt.Errorf("[mid parsing error]: name is empty")
	}

	if mid.Provider == "" {
		return fmt.Errorf("[mid parsing error]: provider is empty")	
	}

	if mid.Namespace == "" {
		return fmt.Errorf("[mid parsing error]: namespace is empty")
	}

	if mid.Version == "" {
		return fmt.Errorf("[mid parsing error]: version is empty")
	}
	return nil
}

func (m MID) String() string {
	if m.Version == "" {
		return fmt.Sprintf("%s/%s/%s", m.Namespace, m.Provider, m.Name)
	}
	
	return fmt.Sprintf("%s/%s/%s/%s", m.Namespace, m.Provider, m.Name, m.Version)
}

func (m MID) Path() string {
	if m.Version == "" {
		fmt.Println("version is empty")
		return ""
	}
	return fmt.Sprintf("%s/%s/%s/%s", m.Namespace, m.Provider, m.Name, m.Version)
}

func (m MID) Filename() string {
	return fmt.Sprintf("%s.zip",m.Name)
}

func (m MID) Filepath() string {
	if m.Version == "" {
		fmt.Println("version is empty")
		return ""
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s.zip",m.Namespace,m.Provider,m.Name,m.Version,m.Name)
}

