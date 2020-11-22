package domain

const (
	successStatus = "success"
	errorStatus   = "error"

	UnExpectedError = "Unexpected error"
)

type Response struct {
	Status string      `json:"status"`
	Result interface{} `json:"result,omitempty"`
	Error  *Error      `json:"error,omitempty"`
}

type Error struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// SuccessResponse simplify send a successfully response
func SuccessResponse(result interface{}) Response {
	return Response{
		Status: successStatus,
		Result: result,
		Error:  nil,
	}
}

// ErrorResponse simplify send an error response
func ErrorResponse(error Error) Response {
	return Response{
		Status: errorStatus,
		Result: nil,
		Error:  &error,
	}
}
