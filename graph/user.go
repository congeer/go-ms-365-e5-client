package graph

import (
	"encoding/json"
	"go-ms-365-e5-sdk/graph/response"
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

func (r *UserRequest) Info() (*response.UserInfo, error) {
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
	resp := response.UserInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (r UserRequest) DriveList() ([]response.Drive, error) {
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
	resp := response.DriveListResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Value, nil
}
