package storage

import "errors"

var (
	ErrDublicateUser = errors.New("dublicate user")
	ErrUserNotFound  = errors.New("user not found")
)
