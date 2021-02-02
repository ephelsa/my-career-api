package domain

const (
	SuccessStatus = "success"
	ErrorStatus   = "error"

	UnexpectedError = "Unexpected error"
	ResourceEmpty   = "Resource is empty"
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
		Status: SuccessStatus,
		Result: result,
		Error:  nil,
	}
}

// ErrorResponse simplify send an error response
func ErrorResponse(error Error) Response {
	return Response{
		Status: ErrorStatus,
		Result: nil,
		Error:  &error,
	}
}
