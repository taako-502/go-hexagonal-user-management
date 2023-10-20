package user_service

import (
	"go-sample-api/application/domain"
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T){
	repository := user_secondary_adapter.NewFakeUserRepository()
	user := &domain.User {
		Username:  "hogepiyo",
		Email:     "test@test.com",
	}

	require.NoError(t, Create(repository, user))
}