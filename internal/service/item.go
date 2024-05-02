package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type ItemService struct {
	repo repository.TodoItem
	listRepo repository.TodoList
}

func NewItemService(repo repository.TodoItem, listRepo repository.TodoList) *ItemService {
	return &ItemService{repo: repo, listRepo: listRepo}
}

func(s *ItemService) Create(userId, listId int, item models.TodoItem) (int, error){
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func(s *ItemService) GetAll(userId, listId int) ([]models.TodoItem, error){
	return s.repo.GetAll(userId, listId);
}

func(s *ItemService) GetById(userId, itemId int) (models.TodoItem, error){
	return s.repo.GetById(userId, itemId)
}

func(s *ItemService) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func(s *ItemService) Update(userId, itemId int, input models.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}