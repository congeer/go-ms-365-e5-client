package auth

import (
	"net/http"
	"net/url"
	"time"
)

// OAuthHandler handle auth path
func OAuthHandler(cli *Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handleRedirectUrl(w, r)
		authUrl := cli.AuthCodeURL()
		http.Redirect(w, r, authUrl, 302)
	}
}

func handleRedirectUrl(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)
	if queryParts["redirectUrl"] != nil {
		http.SetCookie(w, &http.Cookie{
			Name:     "redirectUrl",
			Value:    queryParts["redirectUrl"][0],
			Expires:  time.Now().Add(time.Duration(600) * time.Second),
			HttpOnly: true,
		})
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "redirectUrl",
			Value:    "",
			MaxAge:   -1,
			Expires:  time.Now().Add(time.Duration(-600) * time.Second),
			HttpOnly: true,
		})
	}
}

// CallbackHandler handle callback path
func CallbackHandler(cli *Client, tokenHandle func(w http.ResponseWriter, r *http.Request, token *Token, err error)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParts, _ := url.ParseQuery(r.URL.RawQuery)
		code := queryParts["code"][0]
		state := queryParts["state"][0]
		token, err := cli.GetToken(code, state)
		if tokenHandle != nil {
			tokenHandle(w, r, token, err)
		}
	}
}
