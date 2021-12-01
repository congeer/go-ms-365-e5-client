package auth

import (
	"fmt"
	"net/http"
	"net/url"
)

// OAuthHandler handle auth path
func OAuthHandler(cli *OAuthClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authUrl := cli.AuthCodeURL()
		http.Redirect(w, r, authUrl, 302)
	}
}

// CallbackHandler handle callback path
func CallbackHandler(cli *OAuthClient) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParts, _ := url.ParseQuery(r.URL.RawQuery)
		code := queryParts["code"][0]
		state := queryParts["state"][0]
		token, err := cli.GetToken(code, state)
		fmt.Println(token, err)
		refreshToken, err := cli.RefreshToken(token)
		fmt.Println(refreshToken, err)
	}
}
