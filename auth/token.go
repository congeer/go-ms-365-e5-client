package auth

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
	"time"
)

type Token struct {
	*oauth2.Token
	cli *Client
	ctx context.Context

	Scope        []string
	ExpiresIn    float64
	ExtExpiresIn float64
	ExtExpiry    time.Time
}

func (t Token) HttpClient() *http.Client {
	return t.cli.HttpClient(t.ctx, t)
}

func (t *Token) Refresh() error {
	source := t.cli.conf.TokenSource(t.ctx, t.Token)
	newSource, err := source.Token()

	if err != nil {
		return err
	}

	if newSource.AccessToken == t.AccessToken {
		return nil
	}

	t.Scope = strings.Split(newSource.Extra("scope").(string), " ")
	t.ExpiresIn = newSource.Extra("expires_in").(float64)
	t.ExtExpiresIn = newSource.Extra("ext_expires_in").(float64)
	if v := t.ExtExpiresIn; v != 0 {
		t.ExtExpiry = time.Now().Add(time.Duration(v) * time.Second)
	}
	t.Token = newSource
	return nil
}
