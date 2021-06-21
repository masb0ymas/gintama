package helpers

type response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
}

func BuildResponse(message string, code int, data interface{}, total int64) response {
	jsonResponse := response{
		Message: message,
		Code:    code,
		Data:    data,
		Total:   total,
	}

	return jsonResponse
}
