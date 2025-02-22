package user_service

import (
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService()
	user := &user_model.User{
		Username: "hogepiyo",
		Email:    "test@test.com",
	}
	require.NoError(t, u.Update(repository, user))
}
