package hclconfig

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsimple"
)

const (
	ENV_PATH = "INPUT_CONFIG_PATH"
	DEFAULT_CONFIG_PATH = "terrapak.hcl"
	ENV_TERRAPAK_KEY = "INPUT_TERRAPAK_KEY"
)

var (
	defaultConfig     = &Config{}
)

type Config struct {
	Terrapak TerrapakConfig `hcl:"terrapak,block"`
	Modules []ModuleConfig `hcl:"module,block"`
	Remain  hcl.Body `hcl:",remain"`
}

type TerrapakConfig struct {
	Hostname string `hcl:"hostname"`
	Namespace string `hcl:"organization,optional"`
}

type ModuleConfig struct {
	Name      string `hcl:"name,label"`
	Path      string `hcl:"path"`
	Provider  string `hcl:"provider"`
	Namespace string `hcl:"organization,optional"`
	Version   string `hcl:"version"`
}

func (mc *ModuleConfig) GetNamespace(namespace string) string {
	gc := Default()
	if namespace == "" {
		return gc.Terrapak.Namespace
	}
	return namespace
}

func Default() (config *Config) {
	def := defaultConfig
	return def
}

func Load(configPath *string) (config *Config, err error) {
	config = &Config{}

	err = hclsimple.DecodeFile(*configPath,nil,config); if err != nil {
		return nil, fmt.Errorf("error decoding file: %w",err)
	}
	valid, msg := isValid(config); if !valid {
		return nil, fmt.Errorf(msg)
	}

	config.Terrapak.Hostname = fmt.Sprintf("https://%s",config.Terrapak.Hostname)
	
	defaultConfig = config
	return config, nil
}

func isValid(config *Config) (bool,string) {

	if len(config.Modules) == 0 {
		return false,"[ERROR] - at least one module is required in terrapak.hcl"
	}

	for _, module := range config.Modules {
		if module.Name == "" {
			return false,"[ERROR] - name is required in module block"
		}
		if module.Path == "" {
			return false,"[ERROR] - path is required in module block"
		}

		if module.Provider == "" {
			return false,"[ERROR] - provider is required in module block"
		}

		if module.Version == "" {
			return false,"[ERROR] - version is required in module block"
		}

		if module.Namespace == "" && config.Terrapak.Namespace == "" {
			return false,"[ERROR] - namespace is required in module block or terrapak block"
		}

		ver, err := semver.NewVersion(module.Version); if err != nil {
			return false,"[ERROR] - invalid version number"
		}

		if ver != nil {
			if ver.Major() == 0 && ver.Minor() == 0 && ver.Patch() == 0 {
				return false,"[ERROR] - version cannot be 0.0.0"
			}
		}
	}



	return true,""
}