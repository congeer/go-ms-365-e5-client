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

type DriveItemUpdateRequest struct {
	ParentReference base.ParentReference `json:"parentReference"`
	Name            string               `json:"name,omitempty"`
}

func NewDriveItemUpdateRequest(pathId, rename string) DriveItemUpdateRequest {
	return DriveItemUpdateRequest{
		ParentReference: base.ParentReference{
			Id: pathId,
		},
		Name: rename,
	}
}

func NewDriveItemCopyRequest(driveId, pathId, name string) DriveItemUpdateRequest {
	return DriveItemUpdateRequest{
		ParentReference: base.ParentReference{
			DriveId: driveId,
			Id:      pathId,
		},
		Name: name,
	}
}
