package graph

import (
	"encoding/json"
	"fmt"
	"go-ms-365-e5-sdk/auth"
	"go-ms-365-e5-sdk/graph/response"
	"net/http"
	"strings"
)

var BaseUrl = "https://graph.microsoft.com/v1.0"

type Request struct {
	token *auth.Token
	path  []string
}

func NewRequest(token *auth.Token) *Request {
	return &Request{
		token: token,
		path:  []string{BaseUrl},
	}
}

func (r *Request) client() *http.Client {
	return r.token.HttpClient()
}

func (r *Request) AppendPath(path string) {
	r.path = append(r.path, path)
}

func (r *Request) Get() (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	resp, err := client.Get(url)
	return resp, err
}

func (r *Request) Users(userId string) *UserRequest {
	r.AppendPath("users")
	r.AppendPath(userId)
	return &UserRequest{r}
}

func (r *Request) Me() *UserRequest {
	r.AppendPath("me")
	return &UserRequest{r}
}

func (r *Request) DriveById(id string) *DriveRequest {
	r.AppendPath("drives")
	r.AppendPath(id)
	return &DriveRequest{req: r}
}

func handlerError(body []byte, status int) error {
	errResp := response.ErrorResponse{}
	err := json.Unmarshal(body, &errResp)
	if err != nil {
		return err
	}
	errResp.Status = status
	return errResp
}
