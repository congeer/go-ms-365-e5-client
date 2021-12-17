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
