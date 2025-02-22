package user_primary_adapter

import (
	user_service "go-hexagonal-user-management/core/services/user"
	secondary_port "go-hexagonal-user-management/secondary/port"
	"net/http"
	"strconv"
)

func Delete(u user_service.UserService, a secondary_port.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "ID must be an integer", http.StatusBadRequest)
			return
		}
		if err := u.Delete(a, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
