package repository

import (
	todo "KVANTAKT_PlanNyam"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type JsonChequePostgres struct {
	db *sqlx.DB
}

func NewJsonChequePostgres(db *sqlx.DB) *JsonChequePostgres {
	return &JsonChequePostgres{db: db}
}

func (r *JsonChequePostgres) ParseJsonCheque(listId int, item []byte) ([]int, error) {
	var resultId []int
	var resultParse []todo.JsonCheque
	err := json.Unmarshal(item, &resultParse)
	if err != nil {
		log.Fatal(err)
	}
	resultJson := resultParse[0].Ticket.Document.Receipt.Item
	for _, item := range resultJson {
		tx, err := r.db.Begin()
		if err != nil {
			return nil, err
		}
		var itemId int
		createItemQuery := fmt.Sprintf("INSERT INTO %s (title, quantity, price, description) VALUES ($1, $2, $3, $4) RETURNING id", todoItemsTable)
		row := tx.QueryRow(createItemQuery, item.Name, item.Quantity, item.Price, "")
		err = row.Scan(&itemId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
		_, err = tx.Exec(createListItemsQuery, listId, itemId)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		resultId = append(resultId, itemId)
		tx.Commit()
	}

	return resultId, nil
}
