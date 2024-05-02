package service

import (
	"todo-list/internal/models"
	"todo-list/internal/repository"
)

type ItemService struct {
	repo repository.TodoItem
}

func NewItemService(repo repository.TodoItem) *ItemService {
	return &ItemService{repo: repo}
}

func(s *ItemService) Create(userId, listId int, item models.TodoItem) (int, error){
	return 0, nil;
}

func(s *ItemService) GetAll(userId, listId int) ([]models.TodoItem, error){
	return nil, nil;
}

func(s *ItemService) GetById(userId, itemId int) (models.TodoItem, error){
	return models.TodoItem{}, nil;
}

func(s *ItemService) Delete(userId, itemId int) error {
	return nil
}

func(s *ItemService) Update(userId, itemId int, input models.UpdateItemInput) error {
	return nil
}