package handler

import "net/http"

type Handler struct {
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/sign-up", h.SignUp)
	mux.HandleFunc("/sign-in", h.SignIn)

	return mux
}
