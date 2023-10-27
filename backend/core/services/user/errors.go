package user_service

import "errors"

var UserNotFoundError = errors.New("user not found error")
var UserDuplicateError = errors.New("user duplicate error")
