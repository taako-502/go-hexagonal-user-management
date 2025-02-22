package user_primary_adapter

import (
	"encoding/json"
	"errors"
	user_model "go-hexagonal-user-management/core/models"
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
)

func (a *UserPrimaryAdapter) Create(u user_service.UserService, ur secondary_port.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request UserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := a.validate.Struct(request); err != nil {
			http.Error(w, "validation error", http.StatusBadRequest)
			return
		}
		user, err := user_model.NewUser(request.Username, request.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err = u.Create(ur, user)
		if err != nil {
			if errors.Is(err, user_service.ErrUserDuplicate) {
				http.Error(w, err.Error(), http.StatusConflict)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		response := &UserResponse{
			Id:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	})
}
