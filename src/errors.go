package instagram

import "errors"

var (
	ErrConfigRequired   = errors.New("config required")
	ErrUsernameRequired = errors.New("username required")
	ErrPasswordRequired = errors.New("password required")
	ErrRequestRequired  = errors.New("request required")
)
