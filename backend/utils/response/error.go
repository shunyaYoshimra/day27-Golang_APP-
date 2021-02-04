package response

import "net/http"

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func BadRequest(message string) Error {
	response := Error{
		Status:  http.StatusBadRequest,
		Message: message,
	}
	return response
}

func NotFound(message string) Error {
	response := Error{
		Status:  http.StatusNotFound,
		Message: message,
	}
	return response
}

func Conflict(message string) Error {
	response := Error{
		Status:  http.StatusConflict,
		Message: message,
	}
	return response
}
