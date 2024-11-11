package repos

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/Kennedy-lsd/ExpenseTracker/data"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll(category string) ([]data.Purchase, error) {
	var purchases []data.Purchase
	var query string
	var args []interface{}

	if category != "" {
		query = "SELECT * FROM purchases WHERE category = $1 ORDER BY id ASC"
		args = append(args, category)
	} else {
		query = "SELECT * FROM purchases"
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var purchase data.Purchase

		err = rows.Scan(&purchase.ID, &purchase.Title, &purchase.Price, &purchase.Date, &purchase.Category)
		if err != nil {
			return nil, err
		}

		purchases = append(purchases, purchase)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return purchases, nil
}

func (r *Repository) Create(purchase *data.SetPurchase) error {
	query := "INSERT INTO purchases (title, price, category) VALUES ($1, $2, $3) RETURNING id, date"

	err := r.DB.QueryRow(query, purchase.Title, purchase.Price, strings.ToLower(purchase.Category)).Scan(&purchase.ID, &purchase.Date)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(id int64) error {
	query := "DELETE FROM purchases WHERE id = $1"

	result, err := r.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error checking affected rows: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no stocks found with the given ID")
	}

	return nil
}
