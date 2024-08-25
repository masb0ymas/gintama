package utils

type successResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type failureResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func SuccessResponse(code int32, message string, data interface{}) *successResponse {
	return &successResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func FailureResponse(code int32, message string, errors interface{}) *failureResponse {
	return &failureResponse{
		Code:    code,
		Message: message,
		Errors:  errors,
	}
}
