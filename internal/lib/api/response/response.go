package response

const (
	StatusOk        = "Ok"
	StatusError     = "Error"
	StatusNoContent = "No Content"
)

func OK() Response {
	return Response{
		Status: StatusOk,
	}
}

func NoContent() Response {
	return Response{
		Status: StatusNoContent,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}

// default response
type Response struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
