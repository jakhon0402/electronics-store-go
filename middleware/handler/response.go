package handler

type Response struct {
	StatusCode int
	Data       interface{}
	Err        *ErrorResponse
}

func NewSuccessResponse(status int, data interface{}) *Response {
	return &Response{
		StatusCode: status,
		Data:       data,
	}
}

func NewErrorResponse(status int, code ErrorCode, message string, details interface{}) *Response {
	return &Response{
		StatusCode: status,
		Err: &ErrorResponse{
			Code:    code,
			Message: message,
			Errors:  details,
		},
	}
}
