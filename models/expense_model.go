package models

type Expense struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}
