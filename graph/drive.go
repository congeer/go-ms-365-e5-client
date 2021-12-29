package graph

import (
	"encoding/json"
	"go-ms-365-e5-sdk/graph/response/drive"
	"io/ioutil"
)

type DriveRequest struct {
	req       *Request
	drivePath []string
}

func (r *DriveRequest) Info() (*drive.Drive, error) {
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
	resp := drive.Drive{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *DriveRequest) Root() (*drive.Item, error) {
	return r.Item("root")
}

func (r *DriveRequest) Item(id string) (*drive.Item, error) {
	r.req.AppendPath("items")
	r.req.AppendPath(id)
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
	resp := drive.Item{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r *DriveRequest) RootChildren() ([]drive.Item, error) {
	return r.ItemChildren("root")
}

func (r *DriveRequest) ItemChildren(id string) ([]drive.Item, error) {
	r.req.AppendPath("items")
	r.req.AppendPath(id)
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
	resp := drive.ItemListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}
