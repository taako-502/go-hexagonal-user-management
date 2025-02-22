package user_service

import (
	user_model "go-hexagonal-user-management/core/models"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserService_FindAll(t *testing.T) {
	u := NewUserService()
	repository := user_secondary_adapter.NewFakeUserRepository()
	want := []user_model.User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}

	t.Run("Success", func(t *testing.T) {
		result, err := u.FindAll(repository)
		require.NoError(t, err)
		require.EqualValues(t, want, result)
	})
}
