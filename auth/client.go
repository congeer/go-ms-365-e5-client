package auth

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
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

type Token struct {
	TokenType    string
	Scope        []string
	ExpiresIn    float64
	ExtExpiresIn float64
	AccessToken  string
	RefreshToken string
}

// OAuthClient use
type OAuthClient struct {
	clientConf *Config
	conf       *oauth2.Config
	mutex      *sync.Mutex
}

// NewClient create client
func NewClient(conf Config) *OAuthClient {
	client := OAuthClient{
		clientConf: &conf,
		mutex:      new(sync.Mutex),
	}
	client.initOAuth2()
	return &client
}

func (c *OAuthClient) initOAuth2() {
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
func (c *OAuthClient) AuthCodeURL() string {
	c.initOAuth2()
	return c.conf.AuthCodeURL(c.clientConf.State, oauth2.AccessTypeOffline)
}

// GetToken converts an authorization code into a token.
func (c *OAuthClient) GetToken(code string, state string) (*Token, error) {
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

	c.mutex.Unlock()
	return token, nil
}

func tokenTrans(sourceToken *oauth2.Token) *Token {
	token := &Token{}
	token.TokenType = sourceToken.Extra("token_type").(string)
	token.Scope = strings.Split(sourceToken.Extra("scope").(string), " ")
	token.ExpiresIn = sourceToken.Extra("expires_in").(float64)
	token.ExtExpiresIn = sourceToken.Extra("ext_expires_in").(float64)
	token.AccessToken = sourceToken.Extra("access_token").(string)
	token.RefreshToken = sourceToken.Extra("refresh_token").(string)
	return token
}

func (c *OAuthClient) RefreshToken(token *Token) (*Token, error) {
	source := c.conf.TokenSource(context.Background(), &oauth2.Token{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       time.Now(),
	})
	newToken, err := source.Token()
	if err != nil {
		return nil, err
	}
	token = tokenTrans(newToken)
	return token, nil
}
