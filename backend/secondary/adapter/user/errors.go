package user_secondary_adapter

import "errors"

var UserNotFoundError = errors.New("There were no records in the user table.")
var UserDuplicateError = errors.New("user duplicate error")
