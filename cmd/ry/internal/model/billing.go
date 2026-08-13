package model

import "time"

type Order struct {
	ID      string    `json:"id"      table:"ID"`
	Product string    `json:"product" table:"PRODUCT"`
	Amount  float64   `json:"amount"  table:"AMOUNT"`
	Status  string    `json:"status"  table:"STATUS"`
	Created time.Time `json:"created" table:"CREATED"`
}
