package exceptions

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}
