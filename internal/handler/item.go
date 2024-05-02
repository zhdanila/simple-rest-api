package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todo-list/internal/models"
)

func(h *Handler) createItem(w http.ResponseWriter, r *http.Request) {
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

	var input models.TodoItem
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listIntId, input)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("id: %d", id)))
}

func(h *Handler) getAllItems(w http.ResponseWriter, r *http.Request) {
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

	items, err := h.services.TodoItem.GetAll(userId, listIntId)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	marshalledItems, err := json.Marshal(&items)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledItems)
}

func(h *Handler) getItemById(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	itemId := r.PathValue("id")
	itemIntId, err := strconv.Atoi(itemId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemIntId)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	marshalledItem, err := json.Marshal(&item)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(marshalledItem)
}

func(h *Handler) updateItem(w http.ResponseWriter, r *http.Request) {
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

	var input models.UpdateItemInput
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.TodoItem.Update(userId, listIntId, input); err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("status: updated"))
}

func(h *Handler) deleteItem(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value("userId").(int)
	if !ok {
		NewErrorResponse(w, http.StatusBadRequest, "user id not found in context")
		return
	}

	itemId := r.PathValue("id")
	itemIntId, err := strconv.Atoi(itemId)
	if err != nil {
		NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.Delete(userId, itemIntId)
	if err != nil {
		NewErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("status:ok")))
}