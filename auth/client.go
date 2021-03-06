package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scope        []string
	State        string
	Tenant       string
}

// Client use
type Client struct {
	clientConf *Config
	conf       *oauth2.Config
	mutex      *sync.Mutex
}

// NewClient create client
func NewClient(conf Config) *Client {
	client := Client{
		clientConf: &conf,
		mutex:      new(sync.Mutex),
	}
	client.initClient()
	return &client
}

func (c *Client) initClient() {
	c.mutex.Lock()
	if c.conf == nil {
		c.conf = &oauth2.Config{
			ClientID:     c.clientConf.ClientID,
			ClientSecret: c.clientConf.ClientSecret,
			RedirectURL:  c.clientConf.RedirectURL,
			Scopes:       c.clientConf.Scope,
			Endpoint:     microsoft.AzureADEndpoint(c.clientConf.Tenant),
		}
	}
	c.mutex.Unlock()
}

// AuthCodeURL returns a URL to Microsoft login page
func (c *Client) AuthCodeURL() string {
	c.initClient()
	return c.conf.AuthCodeURL(c.clientConf.State, oauth2.AccessTypeOffline)
}

// GetToken converts an authorization code into a token.
func (c *Client) GetToken(code string, state string) (*Token, error) {
	c.mutex.Lock()
	if c.conf == nil {
		c.mutex.Unlock()
		return nil, fmt.Errorf("client not init")
	}
	if c.clientConf.State != state {
		c.mutex.Unlock()
		return nil, fmt.Errorf("state is different")
	}
	sourceToken, err := c.conf.Exchange(context.Background(), code, oauth2.AccessTypeOffline)
	if err != nil {
		c.mutex.Unlock()
		return nil, err
	}

	token := tokenTrans(sourceToken)
	token.cli = c

	c.mutex.Unlock()
	return token, nil
}

func tokenTrans(sourceToken *oauth2.Token) *Token {
	token := &Token{
		ctx: context.Background(),
	}
	token.Scope = strings.Split(sourceToken.Extra("scope").(string), " ")
	token.ExpiresIn = sourceToken.Extra("expires_in").(float64)
	token.ExtExpiresIn = sourceToken.Extra("ext_expires_in").(float64)
	if v := token.ExtExpiresIn; v != 0 {
		token.ExtExpiry = time.Now().Add(time.Duration(v) * time.Second)
	}
	token.Token = sourceToken
	return token
}

func (c *Client) GetTokenByRefreshToken(refreshToken string) (*Token, error) {
	c.mutex.Lock()

	source := c.conf.TokenSource(context.Background(), &oauth2.Token{
		RefreshToken: refreshToken,
		Expiry:       time.Now(),
	})
	newToken, err := source.Token()

	if err != nil {
		return nil, err
	}

	token := tokenTrans(newToken)
	token.cli = c

	c.mutex.Unlock()
	return token, nil
}

func (c *Client) HttpClient(ctx context.Context, token Token) *http.Client {
	return c.conf.Client(ctx, token.Token)
}
