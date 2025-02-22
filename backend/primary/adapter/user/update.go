package user_primary_adapter

import (
	"encoding/json"
	"errors"
	user_model "go-hexagonal-user-management/core/models"
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
	"strconv"
)

func (a *UserPrimaryAdapter) Update(u user_service.UserService, ur secondary_port.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "ID must be an integer", http.StatusBadRequest)
			return
		}
		var request UserRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := a.validate.Struct(request); err != nil {
			http.Error(w, "validation error", http.StatusBadRequest)
			return
		}
		user, err := user_model.UpdateUser(id, request.Username, request.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := u.Update(ur, user); err != nil {
			if errors.Is(err, user_service.ErrUserDuplicate) {
				http.Error(w, err.Error(), http.StatusConflict)
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
