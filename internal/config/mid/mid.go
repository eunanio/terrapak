package mid

import (
	"fmt"

	"github.com/Masterminds/semver"
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

	if mid.Version != "" {
		mid.Version = buildVersion(mid.Version)
	}
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

func buildVersion(version string) string {
	safe_version, err := semver.NewVersion(version)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%d.%d.%d", safe_version.Major(), safe_version.Minor(), safe_version.Patch())
}

