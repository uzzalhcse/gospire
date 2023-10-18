package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/uzzalhcse/GoSpire/app/models"
)

var (
	request *Request
)

type Request struct {
	Context *gin.Context
}

func NewRequest(c *gin.Context) *Request {
	return &Request{Context: c}
}
func SetRequest(r *Request) {
	request = r
}

func GetRequest() *Request {
	return request
}

func Context() *gin.Context {
	return request.Context
}

func GetHeader(key string) string {
	return request.Context.GetHeader(key)
}
func SetHeader(key string, value string) {
	request.Context.Header(key, value)
}
func Set(key string, value any) {
	request.Context.Set(key, value)
}
func Get(key string) (value any, exists bool) {
	return request.Context.Get(key)
}
func User() *models.User {
	return request.Context.MustGet("user").(*models.User)
}

func Query() map[string][]string {
	return request.Context.Request.URL.Query()
}

func BindRequestData(c *gin.Context, request interface{}) error {
	contentType := c.Request.Header.Get("Content-Type")

	var err error
	switch contentType {
	case "application/json":
		err = c.ShouldBindJSON(request)
	default:
		err = c.ShouldBind(request)
	}

	if err != nil {
		return err
	}

	return nil
}

func Validate(c *gin.Context, request interface{}) *ValidationError {
	if err := BindRequestData(c, request); err != nil {
		validationError := ValidationError{
			Field:   "",
			Message: "",
			Error:   "Invalid Request Data",
		}
		return &validationError
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		validationError := ValidationError{
			Field:   "",
			Message: "",
			Error:   "Validation Error",
		}

		// Extract individual validation errors
		for _, fieldErr := range err.(validator.ValidationErrors) {
			//validationError := ValidationError{
			//	Field:   fieldErr.Field(),
			//	Message: fieldErr.Tag(),
			//	Error:   fieldErr.Field() + " is " + fieldErr.Tag(),
			//}
			validationError.Field = fieldErr.Field()
			validationError.Message = fieldErr.Tag()
			validationError.Error = fieldErr.Field() + " is " + fieldErr.Tag()
		}

		return &validationError
	}

	return nil
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
type ValidationErrors []ValidationError

func (ve ValidationErrors) Error() string {
	return fmt.Sprintf("%v", ve)
}
