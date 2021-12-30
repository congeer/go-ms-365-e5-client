package response

import (
	"time"
)

type DriveListResponse struct {
	Value []Drive `json:"value"`
}

type Drive struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
	WebUrl      string `json:"webUrl"`
	DriveType   string `json:"driveType"`

	CreatedDateTime      time.Time `json:"createdDateTime"`
	LastModifiedDateTime time.Time `json:"lastModifiedDateTime"`

	CreatedBy      *IdentitySet `json:"createdBy,omitempty"`
	LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`

	Owner *IdentitySet `json:"owner,omitempty"`

	Quota struct {
		Deleted   int64  `json:"deleted"`
		Remaining int64  `json:"remaining"`
		State     string `json:"state"`
		Total     int64  `json:"total"`
		Used      int64  `json:"used"`
	} `json:"quota"`
}

type DriveItemListResponse struct {
	Value []DriveItem `json:"value"`
}

type DriveItem struct {
	Id     string `json:"id"`
	ETag   string `json:"eTag,omitempty"`
	CTag   string `json:"cTag,omitempty"`
	WebUrl string `json:"webUrl"`
	Name   string `json:"name"`
	Size   int64  `json:"size"`

	CreatedBy            *IdentitySet `json:"createdBy,omitempty"`
	LastModifiedBy       *IdentitySet `json:"lastModifiedBy,omitempty"`
	CreatedDateTime      time.Time    `json:"createdDateTime"`
	LastModifiedDateTime time.Time    `json:"lastModifiedDateTime"`

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
