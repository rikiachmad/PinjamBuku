package exceptions

import "errors"

var (
	// ErrInvalidCredentials is thrown when the user credentials are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrInternalServerError is thrown when the server encounters an error
	ErrInternalServerError = errors.New("internal server error")
	// ErrNotFound is thrown when the object is not found
	ErrNotFound = errors.New("not found")
	// ErrBadRequest is thrown when request message is invalid
	ErrBadRequest = errors.New("invalid request message")
	// ErrUnauthorized is thrown when user is unauthorized
	ErrUnauthorized = errors.New("unauthorized")
	// ErrUserAlreadyExists is thrown when the user already exists
	ErrUserAlreadyExists = errors.New("user already exists")
	// ErrCartNotFound is thrown when the cart is not found
	ErrCartNotFound = errors.New("cart not found")
	// ErrCartAlreadyExists is thrown when the cart already exists
	ErrCartAlreadyExists = errors.New("cart already exists")
)
