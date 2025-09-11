package error

import "errors"

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrorPasswordIncorrect  = errors.New("password incorrect")
	ErrUsernameExist        = errors.New("username already exist")
	ErrEmailExist           = errors.New("email already exist")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
)

var UserErrors = []error{
	ErrUserNotFound,
	ErrorPasswordIncorrect,
	ErrUsernameExist,
	ErrPasswordDoesNotMatch,
}
