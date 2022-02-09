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
	resp, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	drive := response.Drive{}
	err = json.Unmarshal(body, &drive)
	if err != nil {
		return nil, err
	}
	return &drive, nil
}

func (r *DriveRequest) Root() *DriveItemRequest {
	return &DriveItemRequest{req: r.req.AppendPath("root")}
}

func (r *DriveRequest) ItemById(id string) *DriveItemRequest {
	return &DriveItemRequest{req: r.req.AppendPath("items").AppendPath(id)}
}

func (r *DriveRequest) ItemByPath(path string) *DriveItemRequest {
	req := r.req.AppendPath("items").AppendPath("root:")
	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}
	if strings.HasSuffix(path, "/") {
		path = path[:len(path)-1]
	}
	return &DriveItemRequest{req: req.AppendPath(path + ":")}
}

func (r *DriveRequest) Delta() ([]response.DriveItem, error) {
	req := r.req.AppendPath("root").AppendPath("delta")
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	list := response.DriveItemListResponse{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return list.Value, nil
}

func (r *DriveRequest) SharedWithMe() ([]response.DriveItem, error) {
	req := r.req.AppendPath("sharedWithMe")
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	list := response.DriveItemListResponse{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return list.Value, nil
}

type DriveItemRequest struct {
	req *Request
}

// Info  Get item
func (r *DriveItemRequest) Info() (*response.DriveItem, error) {
	resp, err := r.req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	item := response.DriveItem{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Children List children
func (r *DriveItemRequest) Children() ([]response.DriveItem, error) {
	req := r.req.AppendPath("children")
	resp, err := req.Get()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	list := response.DriveItemListResponse{}
	err = json.Unmarshal(body, &list)
	if err != nil {
		return nil, err
	}
	return list.Value, nil
}

// Move Rename or move file
func (r *DriveItemRequest) Move(pathId string, rename string) (*response.DriveItem, error) {
	resp, err := r.req.PatchJson(request.NewDriveItemUpdateRequest(pathId, rename))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	item := response.DriveItem{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Rename Only rename file
func (r *DriveItemRequest) Rename(rename string) (*response.DriveItem, error) {
	resp, err := r.req.PatchJson(request.NewDriveItemUpdateRequest("", rename))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	item := response.DriveItem{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Copy File In Path
func (r *DriveItemRequest) Copy(name string) (string, error) {
	req := r.req.AppendPath("copy")
	resp, err := req.PostJson(request.NewDriveItemCopyRequest("", "", name))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 202 {
		return "", handlerError(body, resp.StatusCode)
	}
	location := resp.Header.Get("Location")
	return location, nil
}

// CopyToOtherPath Copy File To Other Path
func (r *DriveItemRequest) CopyToOtherPath(name, pathId string) (string, error) {
	req := r.req.AppendPath("copy")
	resp, err := req.PostJson(request.NewDriveItemCopyRequest("", pathId, name))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 202 {
		return "", handlerError(body, resp.StatusCode)
	}
	location := resp.Header.Get("Location")
	return location, nil
}

// CopyOtherDrive Copy File To Other Drive
func (r *DriveItemRequest) CopyOtherDrive(name, pathId, driveId string) (string, error) {
	req := r.req.AppendPath("copy")
	resp, err := req.PostJson(request.NewDriveItemCopyRequest(driveId, pathId, name))
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 202 {
		return "", handlerError(body, resp.StatusCode)
	}
	location := resp.Header.Get("Location")
	return location, nil
}

// Content Download
func (r *DriveItemRequest) Content() ([]byte, string, error) {
	req := r.req.AppendPath("content")
	resp, err := req.Get()
	if err != nil {
		return nil, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != 200 {
		return nil, "", handlerError(body, resp.StatusCode)
	}
	contentDisposition := resp.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	return body, params["filename"], nil
}

// Upload Upload File
func (r *DriveItemRequest) Upload(reader io.Reader) (*response.DriveItem, error) {
	req := r.req.AppendPath("content")
	resp, err := req.Put(reader, "application/octet-stream")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 && resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	item := response.DriveItem{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *DriveItemRequest) ContentReader() (io.ReadCloser, string, error) {
	req := r.req.AppendPath("content")
	resp, err := req.Get()
	if err != nil {
		return nil, "", err
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, "", err
		}
		return nil, "", handlerError(body, resp.StatusCode)
	}
	contentDisposition := resp.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	return resp.Body, params["filename"], nil
}

// CreateFolder Create folder
func (r *DriveItemRequest) CreateFolder(name string) (*response.DriveItem, error) {
	resp, err := r.req.PostJson(request.NewDefaultCreateFolderRequest(name))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, handlerError(body, resp.StatusCode)
	}
	item := response.DriveItem{}
	err = json.Unmarshal(body, &item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Delete Delete file
func (r *DriveItemRequest) Delete() error {
	resp, err := r.req.Delete()
	if err != nil {
		return err
	}
	if resp.StatusCode != 204 && resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return handlerError(body, resp.StatusCode)
	}
	return nil
}
