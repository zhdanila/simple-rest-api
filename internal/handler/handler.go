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

	mux.HandleFunc("POST /sign-up", h.signUp)
	mux.HandleFunc("POST /sign-in", h.signIn)

	mux.HandleFunc("POST /list/", h.userIdentity(h.createList))
	mux.HandleFunc("GET /list/", h.userIdentity(h.getAllLists))
	mux.HandleFunc("GET /list/{id}", h.userIdentity(h.getListById))
	mux.HandleFunc("DELETE /list/{id}", h.userIdentity(h.deleteList))
	mux.HandleFunc("PUT /list/{id}", h.userIdentity(h.updateList))

	mux.HandleFunc("POST /list/{id}/item/", h.userIdentity(h.createItem))
	mux.HandleFunc("GET /list/{id}/item/", h.userIdentity(h.getAllItems))

	mux.HandleFunc("GET /item/{id}", h.userIdentity(h.getItemById))
	mux.HandleFunc("PUT /item/{id}", h.userIdentity(h.updateItem))
	mux.HandleFunc("DELETE /item/{id}", h.userIdentity(h.deleteItem))

	return mux
}
