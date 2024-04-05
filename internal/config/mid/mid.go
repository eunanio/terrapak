package mid

import (
	"fmt"
	"log/slog"
	"regexp"

	"github.com/Masterminds/semver"
	"github.com/gin-gonic/gin"
)

type MID struct {
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	Namespace 	string `json:"namespace"`
	Version     string `json:"version"`
}

func NewMID(namespace, name, provider, version string) (MID, error) {
	mid := MID{
		Name: name,
		Provider: provider,
		Namespace: namespace,
		Version: version,
	}

	err := validate(mid); if err != nil {
		slog.Error(err.Error())
		return mid , err
	}
	return mid, nil
}

func Parse(c *gin.Context) (mid MID, err error){
	mid.Name = c.Param("name")
	mid.Provider = c.Param("provider")
	mid.Namespace = c.Param("namespace")
	mid.Version = c.Param("version")

	if mid.Version != "" {
		mid.Version, err = buildVersion(mid.Version); if err != nil {
			return mid, err
		}
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

	if !isUrlSafeString(mid.Name) {
		return fmt.Errorf("invalid characters in name")
	}

	if !isUrlSafeString(mid.Provider) {
		return fmt.Errorf("invalid characters in provider")
	}

	if !isUrlSafeString(mid.Namespace) {
		return fmt.Errorf("invalid characters in namespace")
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
		slog.Warn("version is empty")
		return ""
	}
	return fmt.Sprintf("%s/%s/%s/%s", m.Namespace, m.Provider, m.Name, m.Version)
}

func (m MID) Filename() string {
	return fmt.Sprintf("%s.zip",m.Name)
}

func (m MID) Filepath() string {
	if m.Version == "" {
		slog.Warn("version is empty")
		return ""
	}
	return fmt.Sprintf("%s/%s/%s/%s/%s.zip",m.Namespace,m.Provider,m.Name,m.Version,m.Name)
}

func buildVersion(version string) (string, error) {
	safe_version, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d.%d.%d", safe_version.Major(), safe_version.Minor(), safe_version.Patch()),nil
}

func isUrlSafeString(s string) bool {
	reg, err := regexp.Compile("^[a-zA-Z0-9_-]*$")
    if err != nil {
        slog.Error(err.Error())
    }
    return reg.MatchString(s)
}

