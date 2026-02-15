package storage

import "errors"

var (
	ErrUserExists        = errors.New("user has already exist")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassowrd   = errors.New("invalid password")
	ErrURLExists         = errors.New("url exists")
	ErrURLNotFound       = errors.New("url not found")
	ErrAliasNotFound     = errors.New("alias not found")
	ErrIdNotFound        = errors.New("id not found")
	ErrWrongPatchRequest = errors.New("to make complete update use put request")
	ErrUnexpectedRows    = errors.New("unexpected rows affected")
)
