package database

import (
	"fmt"
	"testing"

	"github.com/rumiant348/finance-solver/models"
)

func TestConnection(t *testing.T) {
	db := Connection()
	defer db.Close()

	var expense models.Expense
	err := db.QueryRow("SELECT * FROM expenses LIMIT 1").Scan(&expense.ID, &expense.Category, &expense.Price)
	if err != nil {
		t.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Printf("Got first record from expenses %+v\n", expense)

}
