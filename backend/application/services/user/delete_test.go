package user_service

import (
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T){
	repository := user_secondary_adapter.NewFakeUserRepository()
	require.NoError(t, Delete(repository, 1))
}