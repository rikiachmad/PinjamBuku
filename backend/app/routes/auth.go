package routes

import (
	"github.com/rg-km/final-project-engineering-16/backend/app/handler"

	"github.com/gin-gonic/gin"
)

type RequestHandler struct {
	Gin *gin.Engine
}

// AuthRoutes struct
type AuthRoutes struct {
	handler        RequestHandler
	authController handler.AuthController
}

// Setup user routes
func (s AuthRoutes) Setup() {
	auth := s.handler.Gin.Group("/auth")
	{
		auth.POST("/login", s.authController.SignIn)
		auth.POST("/register", s.authController.Register)
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler RequestHandler,
	authController handler.AuthController,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		authController: authController,
	}
}
