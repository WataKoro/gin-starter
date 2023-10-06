package response

import (
	"net/http"
	"os"
)

type APIResponseList interface {
	GetCode() int
	GetMessage() string
	GetData() interface{}
}

type apiResponseList struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (a apiResponseList) GetCode() int {
	return a.Code
}

func (a apiResponseList) GetMessage() string {
	return a.Message
}

func (a apiResponseList) GetData() interface{} {
	return a.Data
}

func SuccessAPIResponseList(code int, message string, data interface{}) APIResponseList {
	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func ErrorAPIResponse(code int, message string) APIResponseList {
	if os.Getenv("APP_ENV") == "production" {
		switch code {
		case http.StatusBadRequest:
			message = "bad request"
		case http.StatusInternalServerError:
			message = "something wrong on the server. contact server admin."
		}
	}

	return &apiResponseList{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
