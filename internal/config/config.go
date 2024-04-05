package config

import (
	"log/slog"
	"os"
	"strings"
	"terrapak/internal/api/storagesource"

	"log"

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
	ENV_TP_SECRET 		  = "TP_SECRET"
	ENV_TP_GITHUB_ORG 	  = "TP_GITHUB_ORG"
	ENV_TP_AUTH_APP_ID 	  = "TP_AUTH_APP_ID"
	ENV_TP_AUTH_KEY_FILE  = "TP_AUTH_KEY_FILE"
	ENV_CONFIG_FILE 	  = "TP_CONFIG_FILE"
	ENV_STORAGE_PATH 	  = "TP_STORAGE"
	ENV_ORGANIZATION 	  = "TP_ORGANIZATION"
	ENV_TP_USER 		  = "TP_USER"
	ENV_TP_PASSWORD 	  = "TP_PASSWORD"
	ENV_TP_ROLES 		  = "TP_ROLES"
	ENV_REDIS_HOST 		  = "TP_REDIS_HOST"
	ENV_REDIS_PASSWORD 	  = "TP_REDIS_PASSWORD"
)

type Config struct {
	Hostname 	 string `yaml:"hostname"`
	Organization string `yaml:"organization"`
	StoragePath  string `yaml:"storage"`
	Database 	 DatabaseConfig `yaml:"database"`
	Redis 		 RedisConfig `yaml:"redis"`
	AuthProvider AuthProviderConfig `yaml:"auth"`
	StorageSource 	 storagesource.StorageSource
	GitHubAppConfig  GitHubAppConfig `yaml:"github"`
	SecretString string `yaml:"secret"`
}

type AuthProviderConfig struct {
	Type 		 string 	`yaml:"type"`
	RoleByEmail  []string   `yaml:"role_by_email"`
	Organization string 	`yaml:"organization"`
	ClientSecret string 	`yaml:"client_secret"`
	ClientId 	 string 	`yaml:"client_id"`
}

type DatabaseConfig struct {
	Hostname string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type RedisConfig struct {
	Hostname string `yaml:"host"`
	Password string `yaml:"password"`
}

type GitHubAppConfig struct {
	AppId 		 string 	`yaml:"appid"`
	KeyFile 	 string 	`yaml:"keyfile"`
	WebhookSecret string 	`yaml:"webhook_secret"`
	PrivateKey 	 string
}

func Load() Config {
	c := Config{}
	_ , exists  := os.LookupEnv(ENV_CONFIG_FILE)

	if exists {
		contents, err := os.ReadFile(os.Getenv(ENV_CONFIG_FILE))
		if err != nil {
			log.Fatal(err)
		}
		
		err = yaml.Unmarshal(contents, &c); if err != nil {
			log.Fatal(err)
		}
	}

	if c.StoragePath == "" {
		slog.Error("storage path is required, either in config file or env")
	}
	
	storageSource, err := storagesource.NewStorageSource(c.StoragePath); if err != nil {
		log.Fatal(err)
	}
	c.StorageSource = storageSource
	setupEnvs(&c)
	loadKeyFile(&c)
	validate(&c)
	defaultClient = &c

	return c
}

func loadKeyFile(c *Config) {
	if c.GitHubAppConfig.KeyFile != "" {
		contents, err := os.ReadFile(c.GitHubAppConfig.KeyFile); if err != nil {
			log.Fatal(err)
		}
		c.GitHubAppConfig.PrivateKey = string(contents)
	}
}

func validate(c *Config){
	if c.Hostname == "" {
		log.Fatal("hostname is required, either in config file or env")
	}

	if c.Database.Hostname == "" {
		log.Fatal("database hostname is required, either in config file or env")
	}

	if c.Database.Username == "" {
		log.Fatal("database username is required, either in config file or env")
	}

	if c.Database.Password == "" {
		log.Fatal("database password is required, either in config file or env")
	}

	if c.AuthProvider.Type == "" {
		log.Fatal("auth type is required, either in config file or env")
	}

	if c.AuthProvider.Type == "github" && c.AuthProvider.ClientId == "" {
		log.Fatal("auth client id is required for oauth type, please set it in config file or env")
	}

	if c.AuthProvider.Type == "github" && c.AuthProvider.ClientSecret == "" {
		log.Fatal("auth client secret is required for oauth type, please set it in config file or env")
	}

	if c.AuthProvider.Type != "github" && c.AuthProvider.Organization != "" {
		log.Fatal("auth organization is only used for github auth type, please set it in config file or env")
	}

	if c.Organization == "" {
		log.Fatal("organization is required, either in config file or env")
	}

}

func GetDefault() *Config {
	def := defaultClient
	return def
}

func setupEnvs(c *Config){
	_, configFileSet := os.LookupEnv(ENV_CONFIG_FILE)
	val, exists := os.LookupEnv(ENV_DB_HOST)
	
	if exists && !configFileSet {
		c.Database.Hostname = val
	}

	val, exists = os.LookupEnv(ENV_DB_USER)
	
	if exists && !configFileSet {
		c.Database.Username = val
	}

	val, exists = os.LookupEnv(ENV_DB_PASS)
	
	if exists && !configFileSet {
		c.Database.Password = val
	}

	val, exists = os.LookupEnv(ENV_TP_HOST)
	
	if exists && !configFileSet {
		c.Hostname = val
	}

	val, exists = os.LookupEnv(ENV_TP_AUTH_TYPE)

	if exists && !configFileSet {
		c.AuthProvider.Type = val
	}

	val, exists = os.LookupEnv(ENV_TP_AUTH_CLIENT_ID)

	if exists && !configFileSet {
		c.AuthProvider.ClientId = val
	}

	val, exists = os.LookupEnv(ENV_TP_AUTH_SECRET)

	if exists && !configFileSet {
		c.AuthProvider.ClientSecret = val
	}

	val, exists = os.LookupEnv(ENV_STORAGE_PATH)

	if exists && !configFileSet {
		c.StoragePath = val
	}

	val, exists = os.LookupEnv(ENV_TP_ROLES)

	if exists && !configFileSet {
		c.AuthProvider.RoleByEmail = strings.Split(val, ",")
	}

	val, exists = os.LookupEnv(ENV_TP_AUTH_KEY_FILE)
	if exists {
		c.GitHubAppConfig.KeyFile = val
	}


	val, exists = os.LookupEnv(ENV_ORGANIZATION)

	if exists && !configFileSet {
		c.Organization = val
	}

	if c.Organization == "" {
		c.Organization = "Default"
	}
}