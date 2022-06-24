package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
)

func AuthMiddleware(validRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		isAuthorized := false
		for _, role := range validRoles {
			role = strings.ToLower(role)
			if c.Keys[role] == true {
				isAuthorized = true
			}
		}

		if !isAuthorized {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}

func ValidateIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		id := c.Keys["id"]

		reqId := c.Param("id")

		reqIdInt, err := strconv.Atoi(reqId)
		if err != nil {
			presenter.ErrorResponse(c, http.StatusInternalServerError, exceptions.ErrInternalServerError)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		reqIdFloat := float64(reqIdInt)

		if id != reqIdFloat {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
