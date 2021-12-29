package graph

import (
	"encoding/json"
	"go-ms-365-e5-sdk/graph/response/drive"
	"go-ms-365-e5-sdk/graph/response/user"
	"io/ioutil"
)

type UserRequest struct {
	req *Request
}

func (r *UserRequest) DriveDefault() *DriveRequest {
	r.req.AppendPath("drive")
	return &DriveRequest{
		req: r.req,
	}
}

func (r *UserRequest) Info() (*user.Info, error) {
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
	resp := user.Info{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r UserRequest) DriveList() ([]drive.Drive, error) {
	r.req.AppendPath("drives")
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
	resp := drive.DriveListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}
