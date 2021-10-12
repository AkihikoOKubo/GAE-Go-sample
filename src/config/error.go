package config

import (
	"github.com/pkg/errors"
)

var (
	// ErrHTTPClientError is client error
	ErrHTTPClientError = errors.New("http client error")
	// ErrBadRequest is 400 error
	ErrBadRequest = errors.New("400:Bad Request")
	// ErrUnauthorized is 401 error
	ErrUnauthorized = errors.New("401:Unauthorized")
	// ErrForbidden is 403 error
	ErrForbidden = errors.New("403:Forbidden")
	// ErrNotFound is 404 error
	ErrNotFound = errors.New("404:Not Found")
	// ErrMethodNotAllowed is 405 error
	ErrMethodNotAllowed = errors.New("405:Method Not Allowed")
	// ErrMethodNotAllowed is 405 error
	ErrConflicted = errors.New("409:Conflicted")
	// ErrInternalServerError is 500 error
	ErrInternalServerError = errors.New("500:Internal Server Error")
	// ErrBadGateway is 502 error
	ErrBadGateway = errors.New("502:Bad Gateway")
	// ErrServiceUnavailable is 503 error
	ErrServiceUnavailable = errors.New("503:Service Unavailable")
	// ErrGatewayTimeout is 504 error
	ErrGatewayTimeout = errors.New("504:Gateway Timeout")
)