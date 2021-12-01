package auth

import (
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
func CallbackHandler(cli *OAuthClient, tokenHandle func(token *Token), errorHandle func(err error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParts, _ := url.ParseQuery(r.URL.RawQuery)
		code := queryParts["code"][0]
		state := queryParts["state"][0]
		token, err := cli.GetToken(code, state)
		if err != nil {
			errorHandle(err)
		} else if tokenHandle != nil {
			tokenHandle(token)
		}
	}
}
