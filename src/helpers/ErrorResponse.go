package helpers

type errResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorResponse(code int, message string) errResponse {
	jsonResponse := errResponse{
		Code:    code,
		Message: message,
	}

	return jsonResponse
}
