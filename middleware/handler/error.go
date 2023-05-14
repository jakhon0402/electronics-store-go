package handler

type ErrorCode string

const (
	// 400 bad request
	InvalidQueryValue = ErrorCode("InvalidQueryValue")
	InvalidUriValue   = ErrorCode("InvalidUriValue")
	InvalidBodyValue  = ErrorCode("InvalidBodyValue")

	// 404 not found
	NotFoundEntity = ErrorCode("NotFoundEntity")

	// 409 duplicate
	DuplicateEntry = ErrorCode("DuplicateEntry")

	// 500
	InternalServerError = ErrorCode("InternalServerError")
)

type ErrorResponse struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"-"`
}
