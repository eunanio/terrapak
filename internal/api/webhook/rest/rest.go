package rest

import (
	"fmt"
	"net/http"
)

type AuthTransport struct {
    Token string
    Transport http.RoundTripper
}

func New(token string) (client *http.Client) {
	t := &AuthTransport{
		Token: token,
		Transport: http.DefaultTransport,
	}
	client = &http.Client{Transport: t}
	return client
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.Token))
    return t.Transport.RoundTrip(req)
}