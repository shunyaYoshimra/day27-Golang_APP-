package response

import "net/http"

type Success struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string) Success {
	response := Success{
		Status:  http.StatusOK,
		Message: message,
	}
	return response
}
