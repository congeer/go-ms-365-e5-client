package main

import (
	"fmt"
	"go-ms-365-e5-sdk/auth"
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

func main() {
	http.HandleFunc("/auth", auth.OAuthHandler(cli))
	http.HandleFunc("/auth/callback", auth.CallbackHandler(cli, func(token *auth.Token) {
		fmt.Println(token.AccessToken)
		// refresh token and print new Token
		newToken, _ := cli.RefreshToken(token)
		fmt.Println(newToken.AccessToken)
	}, nil))
	http.ListenAndServe(":10000", nil)
}
