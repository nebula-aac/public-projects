package user

import (
	"encoding/json"
	"net/http"

	"github.com/nebula-aac/public-projects/simple-google-wire/internal/domain"
)

type handler struct {
	svc domain.UserService
}

func (h *handler) FindByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		user, err := h.svc.FindByUsername(r.Context(), username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(user)
	}
}
