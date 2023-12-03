package user_secondary_adapter

import "errors"

var ErrUserNotFound = errors.New("there were no records in the user table")
var ErrUserDuplicate = errors.New("user duplicate error")
