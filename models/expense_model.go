package models

type Expense struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Price    float32 `json:"price"`
}

type Amount struct {
	ID    string  `json:"id"`
	Price float32 `json:"price"`
	From  string  `json:"from"`
	To    string  `json:"to"`
}
