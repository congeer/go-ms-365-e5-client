# Microsoft 365 E5 Client for Go

### 权限客户端

权限客户端用来访问微软OAuth2接口，获取token或刷新token

##### 1. 创建客户端

示例代码

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

##### 2. 使用系统默认http服务器

使用预置默认http的处理器的示例

需要传入创建的客户端实例，返回请求的处理器

在callback处理器中传入token的处理方法用来处理token和错误，并且继续返回数据
token需要自己处理保存

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

##### 3. 需要刷新token的时候

使用客户端的RefreshToken方法来刷新token（需要传入旧的token）

```go
    newToken, err := cli.RefreshToken(token)
```
