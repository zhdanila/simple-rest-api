package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-list/internal/models"
)

func(h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusInternalServerError, "user id not found in context")
		return
	}

	var input models.TodoList
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.services.TodoList.Create(id, input)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}

func(h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	lists, err := h.services.TodoList.GetAll(id)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	marshalledLists, err := json.Marshal(&lists)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledLists)
}

func(h *Handler) getListById(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	listId := r.PathValue("id")
	listIntId, err := strconv.Atoi(listId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.GetById(userId, listIntId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	marshalledList, err := json.Marshal(&list)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledList)
}

func(h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	listId := r.PathValue("id")
	listIntId, err := strconv.Atoi(listId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Delete(userId, listIntId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: ok"))
}

func(h *Handler) updateList(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	listId := r.PathValue("id")
	listIntId, err := strconv.Atoi(listId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var input models.UpdateListInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoList.Update(userId, listIntId, input); err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: updated"))
}