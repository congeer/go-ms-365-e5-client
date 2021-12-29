package response

type ErrorResponse struct {
	Status    int `json:"status"`
	ErrorInfo struct {
		Code       string `json:"code"`
		Message    string `json:"message"`
		InnerError struct {
			Date            string `json:"date"`
			RequestId       string `json:"request-id"`
			ClientRequestId string `json:"client-request-id"`
		} `json:"innerError"`
	} `json:"error"`
}

func (r ErrorResponse) Error() string {
	return r.ErrorInfo.Message
}

type Identity struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type IdentitySet struct {
	Application *Identity `json:"application,omitempty"`
	Device      *Identity `json:"device,omitempty"`
	Group       *Identity `json:"group,omitempty"`
	User        *User     `json:"user,omitempty"`
}

type User struct {
	Identity
	Email string `json:"email,omitempty"`
}
