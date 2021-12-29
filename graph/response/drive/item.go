package drive

import (
	"go-ms-365-e5-sdk/graph/response"
	"time"
)

type ItemListResponse struct {
	Value []Item `json:"value"`
}

type Item struct {
	Id     string `json:"id"`
	ETag   string `json:"eTag,omitempty"`
	CTag   string `json:"cTag,omitempty"`
	WebUrl string `json:"webUrl"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`

	CreatedBy            *response.IdentitySet `json:"createdBy,omitempty"`
	LastModifiedBy       *response.IdentitySet `json:"lastModifiedBy,omitempty"`
	CreatedDateTime      time.Time             `json:"createdDateTime"`
	LastModifiedDateTime time.Time             `json:"lastModifiedDateTime"`

	ParentReference struct {
		DriveId   string `json:"driveId"`
		DriveType string `json:"driveType"`
		Id        string `json:"id,omitempty"`
		Path      string `json:"path,omitempty"`
	} `json:"parentReference"`
	FileSystemInfo struct {
		CreatedDateTime      time.Time `json:"createdDateTime"`
		LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`
	} `json:"fileSystemInfo"`

	File *struct {
		MimeType string `json:"mimeType"`
		Hashes   struct {
			QuickXorHash string `json:"quickXorHash"`
		} `json:"hashes"`
	} `json:"file,omitempty"`
	Image *struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"image,omitempty"`
	Folder *struct {
		ChildCount int64 `json:"childCount,omitempty"`
	} `json:"folder,omitempty"`
	Root *struct {
	} `json:"root,omitempty"`

	MicrosoftGraphDownloadUrl string `json:"@microsoft.graph.downloadUrl,omitempty"`
}
