package database

import (
	"fmt"
	"os"
	"testing"
)

func TestConnection(t *testing.T) {
	db := Connection()
	defer db.Close()

	var expense struct {
		id          int
		description string
		amount      float32
	}
	err := db.QueryRow("SELECT * FROM expense LIMIT 1").Scan(&expense.id, &expense.amount, &expense.description,)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Got first record from expenses %+v\n", expense)

}
