package main

import "errors"

var (
	ErrDB                = errors.New("database connection error")
	ErrLoginNotFound     = errors.New("login not found")
	ErrPasswordMissMatch = errors.New("password missmatch")

	// other func errors
	ErrInvalidInput = errors.New("invalid input")
	ErrPwned        = errors.New("password has been pwned")
)
