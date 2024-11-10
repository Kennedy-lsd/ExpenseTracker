package data

import "time"

type Purchase struct {
	ID    uint      `json:"id"`
	Title string    `json:"title"`
	Price string    `json:"price"`
	Date  time.Time `json:"date"`
}

type SetPurchase struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Price string `json:"price"`
}
