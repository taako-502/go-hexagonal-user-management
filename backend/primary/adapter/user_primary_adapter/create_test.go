package user_primary_adapter

import (
	"bytes"
	"encoding/json"
	"go-hexagonal-user-management/core/services/user_service"
	"go-hexagonal-user-management/secondary/adapter/user_secondary_adapter"
	"go-hexagonal-user-management/secondary/secondary_port"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestUserPrimaryAdapter_Create(t *testing.T) {
	u := user_service.NewUserService()
	fake := user_secondary_adapter.NewFakeUserRepository()
	pa := NewUserPrimaryAdapter(validator.New())
	type args struct {
		u  user_service.UserService
		ur secondary_port.UserRepository
	}
	tests := []struct {
		name       string
		args       args
		body       UserRequest
		wantStatus int
	}{
		{
			name:       "Success",
			args:       args{u: u, ur: fake},
			body:       UserRequest{Username: "hogepiyo", Email: "test@test.com"},
			wantStatus: http.StatusCreated,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := pa.Create(tt.args.u, tt.args.ur)
			jsonBody, _ := json.Marshal(tt.body)
			request := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(jsonBody))
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, request)

			if status := response.Code; status != tt.wantStatus {
				t.Errorf("Handler.Add() status = %v, want %v", status, tt.wantStatus)
			}
		})
	}
}
