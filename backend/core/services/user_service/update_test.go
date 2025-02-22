package user_service

import (
	"go-hexagonal-user-management/core/model"
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService()
	user := &model.User{
		Username: "hogepiyo",
		Email:    "test@test.com",
	}
	require.NoError(t, u.Update(repository, user))
}
