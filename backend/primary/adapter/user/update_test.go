package user_primary_adapter

import (
	"bytes"
	"context"
	"encoding/json"
	user_service "go-hexagonal-user-management/core/services/user"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestUserPrimaryAdapter_Update(t *testing.T) {
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
		body       UserRequest
		wantStatus int
	}{
		{
			name:       "Success",
			args:       args{u: u, ur: fake},
			wantStatus: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.body)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPatch, "/"+string(tt.ID), bytes.NewReader(body))
			require.NoError(t, err)

			router := http.NewServeMux()
			router.Handle("PATCH /{id}", a.Update(tt.args.u, tt.args.ur))

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req.WithContext(context.TODO()))

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("UserPrimaryAdapter.Update() status = %v, want %v", status, tt.wantStatus)
			}
		})
	}
}
