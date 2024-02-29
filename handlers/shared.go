package handler

import (
	"imaginai/models"
	"log/slog"
	"net/http"
)

func getAuthenticatedUser(r *http.Request) models.AuthenticatedUser {
	if user, ok := r.Context().Value(models.UserContextKey).(models.AuthenticatedUser); ok {
		return user
	}

	return models.AuthenticatedUser{}
}

func Make(h func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("Internal Server Error", "error", err, "path", r.URL.Path)
		}
	}
}
