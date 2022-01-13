package graph

import (
	"encoding/json"
	"go-ms-365-e5-sdk/graph/request"
	"go-ms-365-e5-sdk/graph/response"
	"io"
	"io/ioutil"
	"mime"
	"strings"
)

type DriveRequest struct {
	req *Request
}

func (r *DriveRequest) Info() (*response.Drive, error) {
	get, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.Drive{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *DriveRequest) Root() *DriveItemRequest {
	r.req.AppendPath("root")
	return &DriveItemRequest{req: r.req}
}

func (r *DriveRequest) ItemById(id string) *DriveItemRequest {
	r.req.AppendPath("items")
	r.req.AppendPath(id)
	return &DriveItemRequest{req: r.req}
}

func (r *DriveRequest) ItemByPath(path string) *DriveItemRequest {
	r.req.AppendPath("items")
	r.req.AppendPath("root:")
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	r.req.AppendPath(path + ":")
	return &DriveItemRequest{req: r.req}
}

func (r *DriveRequest) Delta() ([]response.DriveItem, error) {
	r.req.AppendPath("root")
	r.req.AppendPath("delta")
	get, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItemListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

func (r *DriveRequest) SharedWithMe() ([]response.DriveItem, error) {
	r.req.AppendPath("sharedWithMe")
	get, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItemListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

type DriveItemRequest struct {
	req *Request
}

// Info  Get item
func (r *DriveItemRequest) Info() (*response.DriveItem, error) {
	get, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItem{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Children List children
func (r *DriveItemRequest) Children() ([]response.DriveItem, error) {
	r.req.AppendPath("children")
	get, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItemListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}

// Move Rename or move file
func (r *DriveItemRequest) Move(pathId string, rename string) (*response.DriveItem, error) {
	get, err := r.req.Patch(request.NewDriveItemUpdateRequest(pathId, rename))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItem{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Rename Only rename file
func (r *DriveItemRequest) Rename(rename string) (*response.DriveItem, error) {
	get, err := r.req.Patch(request.NewDriveItemUpdateRequest("", rename))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItem{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Content Download
func (r *DriveItemRequest) Content() ([]byte, string, error) {
	r.req.AppendPath("content")
	get, err := r.req.Get()
	if err != nil {
		return nil, "", err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, "", err
	}
	if get.StatusCode != 200 {
		return nil, "", handlerError(body, get.StatusCode)
	}
	contentDisposition := get.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	return body, params["filename"], nil
}

func (r *DriveItemRequest) ContentReader() (io.ReadCloser, string, error) {
	r.req.AppendPath("content")
	get, err := r.req.Get()
	if err != nil {
		return nil, "", err
	}
	if get.StatusCode != 200 {
		body, err := ioutil.ReadAll(get.Body)
		if err != nil {
			return nil, "", err
		}
		return nil, "", handlerError(body, get.StatusCode)
	}
	contentDisposition := get.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	return get.Body, params["filename"], nil
}

// CreateFolder Create folder
func (r *DriveItemRequest) CreateFolder(name string) (*response.DriveItem, error) {
	get, err := r.req.PostJson(request.NewDefaultCreateFolderRequest(name))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}
	if get.StatusCode != 200 {
		return nil, handlerError(body, get.StatusCode)
	}
	resp := response.DriveItem{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Delete Delete file
func (r *DriveItemRequest) Delete() error {
	get, err := r.req.Delete()
	if err != nil {
		return err
	}
	if get.StatusCode != 204 && get.StatusCode != 200 {
		body, err := ioutil.ReadAll(get.Body)
		if err != nil {
			return err
		}
		return handlerError(body, get.StatusCode)
	}
	return nil
}
