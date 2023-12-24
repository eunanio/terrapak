package discovery

import "github.com/gin-gonic/gin"

type ServiceDescovery struct {
	Modules   string      `json:"modules.v1,omitempty"`
	Login     AuthSchema `json:"login.v1,omitempty"`
}

type AuthSchema struct {
	Client     string   `json:"client"`
	GrantTypes []string `json:"grant_types"`
	Authz      string   `json:"authz"`
	Token      string   `json:"token"`
	Ports      []int    `json:"ports"`
}

func Serve(c *gin.Context) {
	login := AuthSchema{}
	sd := ServiceDescovery{}
	sd.Modules = "/v1/modules"
	login.Client = "terraform-cli"
	login.GrantTypes = []string{"authz_code"}
	login.Authz = "/v1/auth/authorize"
	login.Token = "/v1/auth/token"
	sd.Login = login
	sd.Login.Ports = []int{10000,10010}
	c.JSON(200, sd)
}