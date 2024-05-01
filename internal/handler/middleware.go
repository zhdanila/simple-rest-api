package handler

import (
	"net/http"
	"strings"
)

func (h *Handler) userIdentity(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			NewErrorResponse(w, http.StatusBadRequest, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			NewErrorResponse(w, http.StatusBadRequest, "invalid auth header")
			return
		}
		_, err := h.services.Authorization.ParseToken(headerParts[1])
		if err != nil {
			NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	}
}
