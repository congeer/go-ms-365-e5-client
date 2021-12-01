package graph

import (
	"go-ms-365-e5-sdk/auth"
)

// Client use
type Client struct {
	token *auth.Token
}

// NewClient create client
func NewClient(token *auth.Token) *Client {
	client := Client{
		token: token,
	}
	return &client
}
