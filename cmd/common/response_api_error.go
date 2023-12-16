package common

type ResponseAPIError struct {
	status  string `json:"status"`
	message string `json:"message"`
}
