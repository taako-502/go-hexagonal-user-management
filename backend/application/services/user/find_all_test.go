package user_service

import (
	user_secondary_adapter "go-sample-api/secondary/adapter/user"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAll(t *testing.T){
	t.Run("Success", func(t *testing.T) {
		repository := user_secondary_adapter.NewFakeUserRepository()
		result, err := FindAll(repository)
		require.NoError(t, err)
		require.Len(t, result, 1)
		require.Equal(t, result[0].Username, "test")
		require.Equal(t, result[0].Email, "test@test.com")
	})
}