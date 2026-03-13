package errors

import (
	"errors"
)

var (
    
	ErrAuthenticate = errors.New("Authentiction failed")
    ErrInternalserver = errors.New("Internal Server Error")
	ErrBadrequest = errors.New("Bad Request")

)