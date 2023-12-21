package config

import (
	"fmt"
	"os"
	"terrapak/internal/api/metadata"

	"gopkg.in/yaml.v2"
)

var (
	defaultClient *Config
)

const (
	ENV_DB_HOST 	 	  = "TP_DB_HOST"
	ENV_DB_USER 	 	  = "TP_DB_USER"
	ENV_DB_PASS 	 	  = "TP_DB_PASS"
	ENV_TP_HOST 	 	  = "TP_HOSTNAME"
	ENV_TP_AUTH_TYPE 	  = "TP_AUTH_TYPE"
	ENV_TP_AUTH_CLIENT_ID = "TP_AUTH_CLIENT_ID"
	ENV_TP_AUTH_SECRET 	  = "TP_AUTH_SECRET"
	ENV_TP_BUCKET 		  = "TP_BUCKET"
	ENV_CONFIG_FILE 	  = "TP_CONFIG_FILE"
	ENV_STORAGE_PATH 	  = "TP_STORAGE"
	ENV_ORGANIZATION 	  = "TP_ORGANIZATION"
)

type Config struct {
	Hostname 	 string `yaml:"hostname"`
	BucketName   string `yaml:"bucket"`
	Organization string `yaml:"organization"`
	StoragePath  string `yaml:"storage"`
	Database 	 DatabaseConfig `yaml:"database"`
	AuthProvider AuthProviderConfig `yaml:"auth"`
	StorageSource 	 metadata.StorageSource
}

type AuthProviderConfig struct {
	Type 		 string `yaml:"type"`
	ClientSecret string `yaml:"client_secret"`
	ClientId 	 string `yaml:"client_id"`
}

type DatabaseConfig struct {
	Hostname string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Load() Config {
	c := Config{}
	// DEBUG: remove this for production
	os.Setenv(ENV_CONFIG_FILE, "./config.yml")
	_ , exists  := os.LookupEnv(ENV_CONFIG_FILE)

	if exists {
		contents, err := os.ReadFile(os.Getenv(ENV_CONFIG_FILE))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		
		err = yaml.Unmarshal(contents, &c); if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	storageSource, err := metadata.NewStorageSoruce(c.StoragePath); if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c.StorageSource = storageSource

	setupEnvs(&c)
	validate(&c)
	defaultClient = &c

	return c
}

func validate(c *Config){
	if c.Hostname == "" {
		panic("hostname is required, either in config file or env")
	}

	if c.BucketName == "" {
		panic("bucket name is required, either in config file or env")
	}

	if c.Database.Hostname == "" {
		panic("database hostname is required, either in config file or env")
	}

	if c.Database.Username == "" {
		panic("database username is required, either in config file or env")
	}

	if c.Database.Password == "" {
		panic("database password is required, either in config file or env")
	}

	if c.AuthProvider.Type == "" {
		panic("auth type is required, either in config file or env")
	}

	if c.AuthProvider.Type == "oauth" && c.AuthProvider.ClientId == "" {
		panic("auth client id is required for oauth type, please set it in config file or env")
	}

	if c.AuthProvider.Type == "oauth" && c.AuthProvider.ClientSecret == "" {
		panic("auth client secret is required for oauth type, please set it in config file or env")
	}

	if c.Organization == "" {
		panic("organization is required, either in config file or env")
	}

}

func GetDefault() *Config {
	return defaultClient
}

func setupEnvs(c *Config){

	_, exists := os.LookupEnv(ENV_DB_HOST); if !exists {
		os.Setenv(ENV_DB_HOST, c.Database.Hostname)
	}

	_, exists = os.LookupEnv(ENV_DB_USER); if !exists {
		os.Setenv(ENV_DB_USER, c.Database.Username)
	}

	_, exists = os.LookupEnv(ENV_DB_PASS); if !exists {
		os.Setenv(ENV_DB_PASS, c.Database.Password)
	}

	_, exists = os.LookupEnv(ENV_TP_HOST); if !exists {
		os.Setenv(ENV_TP_HOST, c.Hostname)
	}

	_, exists = os.LookupEnv(ENV_TP_AUTH_TYPE); if !exists {
		os.Setenv(ENV_TP_AUTH_TYPE, c.AuthProvider.Type)
	}

	_, exists = os.LookupEnv(ENV_TP_AUTH_CLIENT_ID); if !exists {
		os.Setenv(ENV_TP_AUTH_CLIENT_ID, c.AuthProvider.ClientId)
	}

	_, exists = os.LookupEnv(ENV_TP_AUTH_SECRET); if !exists {
		os.Setenv(ENV_TP_AUTH_SECRET, c.AuthProvider.ClientSecret)
	}

	_, exists = os.LookupEnv(ENV_TP_BUCKET); if !exists {
		os.Setenv(ENV_TP_BUCKET, c.BucketName)
	}

	_, exists = os.LookupEnv(ENV_STORAGE_PATH); if !exists {
		os.Setenv(ENV_STORAGE_PATH, c.StoragePath)
	}

	_, exists = os.LookupEnv(ENV_ORGANIZATION); if !exists {
		os.Setenv(ENV_ORGANIZATION, c.Organization)
	}

	if c.Organization == "" {
		c.Organization = "Default"
	}

	if c.AuthProvider.Type == "" {
		c.AuthProvider.Type = "PAT"
	}
}