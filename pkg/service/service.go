package service

import (
	todo "KVANTAKT_PlanNyam"
	"KVANTAKT_PlanNyam/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type JsonCheque interface {
	ParseJsonCheque(userId int, item []byte) ([]int, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	JsonCheque
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		JsonCheque:    NewJsonChequeService(repos.JsonCheque),
	}
}
