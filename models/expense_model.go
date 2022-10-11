package models

type Expense struct {
	ID	int	`json:"id"`
	Description	string	`json:"description"`
	Amount	float64 `json:"amount"`
}