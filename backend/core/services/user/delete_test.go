package user_service

import (
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T){
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService(nil)
	require.NoError(t, u.Delete(repository, 1))
}