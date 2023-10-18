package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, webResponse{
		Code:    http.StatusOK,
		Status:  http.StatusText(http.StatusOK),
		Message: message,
		Data:    data,
	})
}
func Error(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, webResponse{
		Code:    http.StatusBadRequest,
		Status:  http.StatusText(http.StatusBadRequest),
		Data:    data,
		Message: message,
	})
}
func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, webResponse{
		Code:    http.StatusBadRequest,
		Status:  http.StatusText(http.StatusBadRequest),
		Message: message,
	})
}

func UnAuthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, webResponse{
		Code:    http.StatusUnauthorized,
		Status:  http.StatusText(http.StatusUnauthorized),
		Message: message,
	})
}
func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, webResponse{
		Code:    http.StatusForbidden,
		Status:  http.StatusText(http.StatusForbidden),
		Message: message,
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, webResponse{
		Code:    http.StatusNotFound,
		Status:  http.StatusText(http.StatusNotFound),
		Message: message,
	})
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, webResponse{
		Code:    http.StatusInternalServerError,
		Status:  http.StatusText(http.StatusInternalServerError),
		Message: "Internal Server Error",
	})
}
