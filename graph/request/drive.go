package request

import "go-ms-365-e5-sdk/graph/base"

type CreateFolderRequest struct {
	Name   string `json:"name"`
	Folder struct {
	} `json:"folder"`
	ConflictBehavior string `json:"@microsoft.graph.conflictBehavior"`
}

func NewDefaultCreateFolderRequest(name string) CreateFolderRequest {
	return CreateFolderRequest{
		Name:             name,
		ConflictBehavior: "rename",
	}
}

type DriveItemMoveRequest struct {
	ParentReference base.ParentReference `json:"parentReference"`
	Name            string               `json:"name,omitempty"`
}

func NewDriveItemMoveRequest(pathId, rename string) DriveItemMoveRequest {
	return DriveItemMoveRequest{
		ParentReference: base.ParentReference{
			Id: pathId,
		},
		Name: rename,
	}
}
