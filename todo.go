package KVANTAKT_PlanNyam

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Quantity    string `json:"quantity" db:"quantity" binding:"required"`
	Price       string `json:"price" db:"price" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"Description"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil {
		return errors.New("Update structure has no values")
	}
	return nil
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Quantity    *string `json:"quantity"`
	Price       *string `json:"price"`
	Description *string `json:"Description"`
	Done        *bool   `json:"done" db:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Done == nil {
		return errors.New("Update structure has no values")
	}
	return nil
}

type JsonCheque struct {
	Ticket struct {
		Document struct {
			Receipt struct {
				Item []struct {
					Name     string  `json:"name"`
					Quantity float32 `json:"quantity"`
					Price    float32 `json:"price"`
				} `json:"items"`
			} `json:"receipt"`
		} `json:"document"`
	} `json:"ticket"`
}
