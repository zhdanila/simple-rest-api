package handler

import (
	"fmt"
	"net/http"
)

func(h *Handler) CreateList(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusInternalServerError, "user id not found in context")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}