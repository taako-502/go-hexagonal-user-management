package user_primary_adapter

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"

	user_service "go-hexagonal-user-management/core/services/user"
	user_secondary_adapter "go-hexagonal-user-management/secondary/adapter/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
)

func TestUserPrimaryAdapter_FindAll(t *testing.T) {
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
		want       []UserResponse
		wantStatus int
	}{
		{
			name: "Success",
			args: args{u: u, ur: fake},
			want: []UserResponse{
				{ID: 1, Username: "user1", Email: "user1@example.com"},
				{ID: 2, Username: "user2", Email: "user2@example.com"},
			},
			wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/users", nil)
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := a.FindAll(tt.args.u, tt.args.ur)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantStatus {
				t.Errorf("status code got %v, want %v", status, tt.wantStatus)
			}

			var actual []UserResponse
			if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			// 期待値と実際のレスポンスを比較
			if !reflect.DeepEqual(actual, tt.want) {
				t.Errorf("unexpected response: got %+v, want %+v", actual, tt.want)
			}
		})
	}
}
