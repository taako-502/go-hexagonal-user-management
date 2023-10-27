package user_service

import (
	"go-hexagonal-user-management/core/domain"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T){
	repository := user_secondary_adapter.NewFakeUserRepository()
	u := NewUserService(nil)
	user := &domain.User {
		Username:  "hogepiyo",
		Email:     "test@test.com",
	}
	require.NoError(t, u.Create(repository, user))
}
