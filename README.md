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
    http.HandleFunc("/auth", auth.OAuthHandler(cli))
    http.HandleFunc("/auth/callback", auth.CallbackHandler(cli, func(token *auth.Token) {
        // handle token
        fmt.Println(token.AccessToken)
    }, nil))
    http.ListenAndServe(":10000", nil)
```

##### 3. When Need Refresh Token

```go
    newToken, err := cli.RefreshToken(token)
```
