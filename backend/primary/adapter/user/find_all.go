package user_primary_adapter

import (
	"encoding/json"
	"errors"
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
)

func (a *UserPrimaryAdapter) FindAll(u user_service.UserService, ur secondary_port.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := u.FindAll(ur)
		if err != nil {
			if errors.Is(err, user_service.ErrUserNotFound) {
				http.Error(w, err.Error(), http.StatusNotFound)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		var responses []UserResponse
		for _, user := range users {
			responses = append(responses, UserResponse{
				Id:       user.Id,
				Username: user.Username,
				Email:    user.Email,
			})
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(responses); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	})
}
