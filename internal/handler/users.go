package handler

import (
	"encoding/json"
	"net/http"
)

func(h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	people, err := h.services.Users.GetAll()
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	marshalledPeople, err := json.Marshal(people)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledPeople)
}

