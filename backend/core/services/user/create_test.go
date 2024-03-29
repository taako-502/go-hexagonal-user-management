package user_service

import (
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService(nil)
	user := &user_model.User{
		Username: "hogepiyo",
		Email:    "test@test.com",
	}
	user, err := u.Create(repository, user)
	require.NoError(t, err)
	require.EqualValues(t, 999, user.Id)
}
