package response

type TestResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []TestError `json:"error"`
}

type TestError struct {
	Message string `json:"message"`
	Flag    string `json:"flag"`
}
