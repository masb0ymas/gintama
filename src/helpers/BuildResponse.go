package helpers

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func BuildResponse(code int, message string, data interface{}) response {
	jsonResponse := response{
		Code:    code,
		Message: message,
		Data:    data,
	}

	return jsonResponse
}
