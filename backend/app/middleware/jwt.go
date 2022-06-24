package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rg-km/final-project-engineering-16/backend/app/presenter"
	"github.com/rg-km/final-project-engineering-16/backend/commons/exceptions"
)

//jwt service
type JWTService interface {
	GenerateToken(id int64, email string, isAdmin bool, isLibrary bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	ID      int64  `json:"id"`
	Email   string `json:"email"`
	Admin   bool   `json:"admin"`
	Library bool   `json:"library"`
	jwt.StandardClaims
}

type JwtServices struct {
	secretKey string
	issuer    string
}

func JWTAuthService() JWTService {
	return &JwtServices{
		secretKey: getSecretKey(),
		issuer:    "PinjamBuku",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *JwtServices) GenerateToken(id int64, email string, isAdmin bool, isLibrary bool) string {
	claims := &authCustomClaims{
		id,
		email,
		isAdmin,
		isLibrary,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *JwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, BEARER_SCHEMA) {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			presenter.ErrorResponse(c, http.StatusUnauthorized, exceptions.ErrUnauthorized)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		fmt.Println(claims)
		if len(c.Keys) == 0 {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["library"] = claims["library"]
		c.Keys["admin"] = claims["admin"]
		c.Keys["id"] = claims["id"]

		if claims["library"] == false && claims["admin"] == false {
			c.Keys["user"] = true
		}
	}
}
