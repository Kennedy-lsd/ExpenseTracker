package data

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Purchase struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title" validate:"required,min=3,max=255"`
	Price    float64   `json:"price" validate:"required,numeric"`
	Date     time.Time `json:"date"`
	Category string    `json:"category" validate:"required,oneof=shop food entertainment"`
}

type SetPurchase struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title" validate:"required,min=3,max=255"`
	Price    string    `json:"price" validate:"required,numeric"`
	Date     time.Time `json:"date"`
	Category string    `json:"category" validate:"required,oneof=shop food entertainment"`
}

var validate = validator.New()

func (p *Purchase) Validate() error {
	return validate.Struct(p)
}

func (p *SetPurchase) Validate() error {
	return validate.Struct(p)
}
