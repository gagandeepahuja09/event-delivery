package internal

type error interface {
	Error() string
	Code() int
}

type errorResp struct {
	Error interface{} `json:"error"`
}

type customError struct {
	message string
	code    int
}

func (e customError) Error() string {
	return e.message
}

func (e customError) Code() int {
	return e.code
}

func newCustomError(message string, code int) error {
	return customError{message, code}
}
