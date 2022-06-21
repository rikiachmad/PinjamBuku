package presenter

import (
	"github.com/gin-gonic/gin"
)

type BasicResponse struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	res := BasicResponse{
		Status:  status,
		Success: true,
		Message: "success",
		Data:    data,
	}
	c.JSON(status, res)
}

func ErrorResponse(c *gin.Context, status int, err error) {
	res := BasicResponse{
		Status:  status,
		Success: false,
		Message: err.Error(),
		Data:    nil,
	}
	c.JSON(status, &res)
}
