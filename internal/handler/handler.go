package handler

import (
	"net/http"
	"todo-list/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /sign-up", h.SignUp)
	mux.HandleFunc("POST /sign-in", h.SignIn)

	mux.HandleFunc("GET /get", h.GetAll)

	return mux
}
