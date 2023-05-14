package errors

import "net/http"

type RestErr struct {
	Message     string `json:"messsage"`
	Status      int    `json:"status"`
	Error       string `json:"error"`
	ErrorDetail string `json:"error_detail"`
}

func NewBadRequestError(message, errorDetail string) *RestErr {
	return &RestErr{
		Message:     message,
		Status:      http.StatusBadRequest,
		Error:       "bad_request",
		ErrorDetail: errorDetail,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
