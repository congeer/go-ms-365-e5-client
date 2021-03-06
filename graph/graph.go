package graph

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-ms-365-e5-sdk/auth"
	"go-ms-365-e5-sdk/graph/base"
	"io"
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

func (r Request) AppendPath(path string) *Request {
	r.path = append(r.path, path)
	return &r
}

func (r *Request) Get() (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	resp, err := client.Get(url)
	return resp, err
}

func (r *Request) PostJson(req interface{}) (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(marshal))
	return resp, err
}

func (r *Request) PatchJson(req interface{}) (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	return resp, err
}

func (r *Request) PutJson(req interface{}) (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	return resp, err
}

func (r *Request) Put(reader io.Reader, contentType string) (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	request, err := http.NewRequest(http.MethodPut, url, reader)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", contentType)
	resp, err := client.Do(request)
	return resp, err
}

func (r *Request) Delete() (*http.Response, error) {
	client := r.token.HttpClient()
	url := strings.Join(r.path, "/")
	fmt.Println(url)
	request, err := http.NewRequest(http.MethodDelete, url, bytes.NewBufferString(""))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	return resp, err
}

func (r *Request) Users(userId string) *UserRequest {
	return &UserRequest{r.AppendPath("users").AppendPath(userId)}
}

func (r *Request) Me() *UserRequest {
	return &UserRequest{r.AppendPath("me")}
}

func (r *Request) DriveById(id string) *DriveRequest {
	return &DriveRequest{req: r.AppendPath("drives").AppendPath(id)}
}

func handlerError(body []byte, status int) error {
	errResp := base.ErrorResponse{}
	err := json.Unmarshal(body, &errResp)
	if err != nil {
		return err
	}
	errResp.Status = status
	return errResp
}
