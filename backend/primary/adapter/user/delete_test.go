package user_primary_adapter

import (
	"context"
	"fmt"
	user_service "go-hexagonal-user-management/core/services/user"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestUserPrimaryAdapter_Delete(t *testing.T) {
	a := NewUserPrimaryAdapter(validator.New())
	u := user_service.NewUserService()
	fake := user_secondary_adapter.NewFakeUserRepository()

	type args struct {
		u  user_service.UserService
		ur secondary_port.UserRepository
	}
	tests := []struct {
		name       string
		args       args
		ID         int
		wantStatus int
	}{
		{
			name:       "Success",
			args:       args{u: u, ur: fake},
			wantStatus: http.StatusNoContent,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, "/"+fmt.Sprint(tt.ID), nil)
			require.NoError(t, err)

			router := http.NewServeMux()
			router.Handle("GET /{id}", a.Delete(tt.args.u, tt.args.ur))

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req.WithContext(context.TODO()))

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("UserPrimaryAdapter.Delete() status = %v, want %v", status, tt.wantStatus)
			}
		})
	}
}
