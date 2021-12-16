# Microsoft 365 E5 Client for Go

### Auth Client

Create auth client to connect microsoft 

##### 1. Create Client

```go
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
```

##### 2. Use Http

```go
    func tokenHandler(w http.ResponseWriter, r *http.Request, token *auth.Token, err error) {
        if err != nil {
            io.WriteString(w, err.Error())
            return
        }
        io.WriteString(w, "=====access token====\n")
        io.WriteString(w, token.AccessToken)
        io.WriteString(w, "\n=====refresh token====\n")
        io.WriteString(w, token.RefreshToken)
        io.WriteString(w, "\n=====access token refresh====\n")
        newToken, _ := cli.RefreshToken(token)
        io.WriteString(w, newToken.AccessToken)
    }
	
    http.HandleFunc("/auth", auth.OAuthHandler(cli))
    http.HandleFunc("/auth/callback", auth.CallbackHandler(cli, tokenHanlder))
    http.ListenAndServe(":10000", nil)
```

##### 3. When Need Refresh Token

```go
    newToken, err := cli.RefreshToken(token)
```
