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

func (r *Repository) GetAll(category string) ([]data.Purchase, float64, error) {
	var purchases []data.Purchase
	var totalAmount float64
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
		return nil, 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var purchase data.Purchase

		err = rows.Scan(&purchase.ID, &purchase.Title, &purchase.Price, &purchase.Date, &purchase.Category)
		if err != nil {
			return nil, 0, err
		}

		purchases = append(purchases, purchase)

		totalAmount += purchase.Price
	}

	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return purchases, totalAmount, nil
}

func (r *Repository) Create(purchase *data.SetPurchase) error {
	query := "INSERT INTO purchases (title, price, category) VALUES ($1, $2, $3) RETURNING id, date"

	if err := purchase.Validate(); err != nil {
		return err
	}

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

func (r *Repository) Update(id int64, purchase *data.Purchase) error {
	query := `UPDATE tasks SET title = ?, price = ?, category = ?, date = ? WHERE id = ?`

	if err := purchase.Validate(); err != nil {
		return err
	}

	_, err := r.DB.Exec(query, purchase.Title, purchase.Price, purchase.Category, purchase.Date, id)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	return nil
}
