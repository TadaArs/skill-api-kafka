package errs


type Err struct {
	StatusCode int
	Message string 
}

func (e Err) Error() string {
	return e.Message
}

func NewError(code int , message string) error {
	return Err{
		StatusCode: code, 
		Message: message,
	}
}
