package user_service

import (
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService()
	require.NoError(t, u.Delete(repository, 1))
}
