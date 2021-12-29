package drive

import (
	"go-ms-365-e5-sdk/graph/response"
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

	CreatedBy      *response.IdentitySet `json:"createdBy,omitempty"`
	LastModifiedBy *response.IdentitySet `json:"lastModifiedBy,omitempty"`

	Owner *response.IdentitySet `json:"owner,omitempty"`

	Quota struct {
		Deleted   int64  `json:"deleted"`
		Remaining int64  `json:"remaining"`
		State     string `json:"state"`
		Total     int64  `json:"total"`
		Used      int64  `json:"used"`
	} `json:"quota"`
}
