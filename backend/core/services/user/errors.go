package user_service

import "errors"

var ErrUserNotFound = errors.New("ユーザが見つかりません。")
var ErrUserDuplicate = errors.New("ユーザ名またはメールアドレスが他のユーザと重複しています。")
