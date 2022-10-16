package models

type Expense struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}
