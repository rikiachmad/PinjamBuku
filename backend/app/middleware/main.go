package middleware

import (
	"net/http"
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
