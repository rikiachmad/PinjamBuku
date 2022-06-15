package presenter

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, status int, data interface{}) error {
	res := gin.H{
		"status":  status,
		"success": true,
		"message": "Success",
		"data":    data,
	}
	return c.JSON(status, &res)
}

func ErrorResponse(c *gin.Context, status int, err error) error {
	res := gin.H{
		"status":  status,
		"success": false,
		"message": err.Error(),
		"data":    nil,
	}
	return c.JSON(status, &res)
}