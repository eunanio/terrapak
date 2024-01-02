package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
	"terrapak/internal/api/auth/providers/github"
	"terrapak/internal/api/auth/providers/types"
	"terrapak/internal/api/auth/roles"
	"terrapak/internal/config"
	"terrapak/internal/db/entity"
	"terrapak/internal/db/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

var (
	codeVerifier = generateCodeVerifier()
	codeChallenge = generateCodeChallenge(codeVerifier)
)

type AuthProvider interface{
	Name() string
	Config() (conf oauth2.Config)
	Authenticate(token string)
	UserInfo(token string) (types.UserInfo,error)
}

type OAuthToken struct {
	AccessToken string `json:"access_token"`
}

func GetAuthProvider() AuthProvider {
	gc := config.GetDefault()
	switch gc.AuthProvider.Type {
		case "github":
			github := github.New()
			return github
	}
	return nil

}


func Authorize(c *gin.Context) {
	//..
	//sessions := sessions.Default(c)
	gc := config.GetDefault()
	state := uuid.New().String()
	// sessions.Set("state", state)
	redirect := fmt.Sprintf("https://%s/v1/auth/callback", gc.Hostname)
	provider := GetAuthProvider()

	conf := provider.Config()
	conf.RedirectURL = redirect
	url := conf.AuthCodeURL(state,oauth2.SetAuthURLParam("code_challenge", codeChallenge),oauth2.SetAuthURLParam("code_challenge_method", "S256"))
	c.Redirect(302, url)
}

func Token(c *gin.Context) {
	//..
}

func Callback(c *gin.Context) {
	//..
	// sessions := sessions.Default(c)
	// state := sessions.Get("state")
	// fmt.Println(state)
	// if state != c.Query("state") {
	// 	c.JSON(401, gin.H{
	// 		"error": "Invalid state",
	// 	})
	// 	return
	// }
	gc := config.GetDefault()
	provider := GetAuthProvider()
	conf := provider.Config()
	token, err := conf.Exchange(c, c.Query("code"), oauth2.SetAuthURLParam("code_verifier", codeVerifier)); if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	// claims, err := jwt.DecodeJWT(token.AccessToken); if err != nil {
	// 	c.JSON(401, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }
	// fmt.Println(claims)
	syncUserAccounts(token.AccessToken)
	c.Data(200, "text/html; charset=utf-8", []byte(fmt.Sprintf("export TF_TOKEN_%s=%s </br></br>Set this if terraform fails to detect the callback",buildSafeHostname(gc.Hostname), token.AccessToken)))
}

func generateCodeVerifier() string {
    b := make([]byte, 32)
    rand.Read(b)
    return base64.RawURLEncoding.EncodeToString(b)
}

func generateCodeChallenge(verifier string) string {
    s256 := sha256.Sum256([]byte(verifier))
    return base64.RawURLEncoding.EncodeToString(s256[:])
}

func buildSafeHostname(hostname string) string {
	return strings.ReplaceAll(hostname, ".", "_")
}

func syncUserAccounts(access_token string){
	provider := GetAuthProvider()
	us := &services.UserService{}
	info, err := provider.UserInfo(access_token); if err != nil {
		fmt.Println(err)
		return
	 }

	 fmt.Println(info)
	 user := us.FindByExternalID(fmt.Sprintf("%d", info.ID))
	 if user == nil {
		user = &entity.User{}
		user.Email = ""
		user.AuthorityID = fmt.Sprintf("%d", info.ID)
		user.Name = info.Name
		user.Role = roles.Editor
		us.Create(*user)

	 }

	 fmt.Println(user)
	 
}