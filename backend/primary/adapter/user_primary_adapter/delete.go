package user_primary_adapter

import (
	"go-hexagonal-user-management/core/services/user_service"
	"go-hexagonal-user-management/secondary/secondary_port"
	"log"
	"net/http"
	"strconv"
)

func (a *UserPrimaryAdapter) Delete(u user_service.UserService, ur secondary_port.UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "ID must be an integer", http.StatusBadRequest)
			return
		}
		if err := u.Delete(ur, id); err != nil {
			log.Printf("[ERROR] Failed to delete user with ID %d: %v", id, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
