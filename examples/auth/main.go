package main

import (
	"go-ms-365-e5-sdk/auth"
	"io"
	"net/http"
)

var cli = auth.NewClient(auth.Config{
	ClientID:     "client-id",
	ClientSecret: "client-secret",
	RedirectURL:  "http://localhost:10000/auth/callback",
	State:        "test",
	Tenant:       "common",
	Scope: []string{
		"calendars.read",
		"offline_access",
	},
})

var Token *auth.Token

func tokenHandler(w http.ResponseWriter, r *http.Request, token *auth.Token, err error) {
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	cookie, err := r.Cookie("redirectUrl")
	if err == nil {
		redirectUrl := cookie.Value
		if redirectUrl != "" {
			Token = token
			http.Redirect(w, r, redirectUrl, 302)
			return
		}
	}
	io.WriteString(w, "=====access token====\n")
	io.WriteString(w, token.AccessToken)
	io.WriteString(w, "\n=====refresh token====\n")
	io.WriteString(w, token.RefreshToken)
	io.WriteString(w, "\n=====access token refresh====\n")
	newToken, _ := cli.GetTokenByRefreshToken(token.RefreshToken)
	io.WriteString(w, newToken.AccessToken)
	Token = newToken
}

func main() {
	http.HandleFunc("/auth", auth.OAuthHandler(cli))
	http.HandleFunc("/auth/callback", auth.CallbackHandler(cli, tokenHandler))
	http.ListenAndServe(":10000", nil)
}
