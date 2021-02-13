package response

import "net/http"

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Redirect struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Path    string      `json:"path"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string) Success {
	response := Success{
		Status:  http.StatusOK,
		Message: message,
	}
	return response
}

func RedirectResponse(message, path string) Redirect {
	response := Redirect{
		Status:  http.StatusFound,
		Message: message,
		Path:    path,
	}
	return response
}
