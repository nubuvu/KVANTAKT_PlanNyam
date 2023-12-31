package repository

import (
	todo "KVANTAKT_PlanNyam"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todo.User) error
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type JsonCheque interface {
	ParseJsonCheque(userId int, item []byte) ([]int, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
	JsonCheque
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
		JsonCheque:    NewJsonChequePostgres(db),
	}
}
