package graph

import (
	"encoding/json"
	"go-ms-365-e5-sdk/graph/response/drive"
	"io/ioutil"
)

type DriveRequest struct {
	req *Request
}

func (r *DriveRequest) RootChildren() ([]drive.Item, error) {
	r.req.AppendPath("root")
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
