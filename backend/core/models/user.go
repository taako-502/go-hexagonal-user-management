package user_model

import (
	"fmt"
	"strings"

	"errors"
)

type User struct {
	Id       int
	Username string
	Email    string
}

func NewUser(username, email string) (*User, error) {
	if strings.TrimSpace(username) == "" {
		return nil, fmt.Errorf("username cannot be blank, strings.TrimSpace(%s)", username)
	}

	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("email cannot be blank, strings.TrimSpace(%s)", email)
	}

	return &User{
		Username: username,
		Email:    email,
	}, nil
}

func UpdateUser(id int, username, email string) (*User, error) {
	if id <= 0 {
		return nil, errors.New("invalid Id")
	}

	if strings.TrimSpace(username) == "" {
		return nil, fmt.Errorf("username cannot be blank, strings.TrimSpace(%s)", username)
	}

	if strings.TrimSpace(email) == "" {
		return nil, fmt.Errorf("email cannot be blank, strings.TrimSpace(%s)", email)
	}

	return &User{
		Id:       id,
		Username: username,
		Email:    email,
	}, nil
}
