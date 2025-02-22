package user_service

import (
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService()
	user := &model.User{
		Username: "hogepiyo",
		Email:    "test@test.com",
	}
	user, err := u.Create(repository, user)
	require.NoError(t, err)
	require.EqualValues(t, 999, user.ID)
}
