package repository

import (
	"database/sql"
	"github.com/rumiant348/finance-solver/models"
	"log"
)

type ExpenseRepository interface {
	CreateTable() error
	GetExpenses() ([]models.Expense, error)
	GetExpenseById(id string) (models.Expense, error)
	CreateExpense(expense models.Expense) (models.Expense, error)
	DeleteExpenseById(id string) (string, error)
}

type expenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository {
	return expenseRepository{db: db}
}

func (er expenseRepository) CreateTable() error {
	er.Ping()
	_, err := er.db.Exec(`CREATE TABLE IF NOT EXISTS expenses(
			id serial PRIMARY KEY, 
    		category VARCHAR (255) NOT NULL, 
    		price FLOAT NOT NULL
	  );`)
	return err
}

func (er expenseRepository) GetExpenses() ([]models.Expense, error) {
	er.Ping()
	rows, err := er.db.Query("SELECT * FROM expenses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var exp models.Expense
		if err := rows.Scan(&exp.ID, &exp.Category, &exp.Price); err != nil {
			return expenses, err
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

func (er expenseRepository) CreateExpense(expense models.Expense) (models.Expense, error) {
	er.Ping()
	err := er.db.QueryRow("INSERT INTO expenses(category, price) VALUES($1, $2)  RETURNING id",
		expense.Category, expense.Price).Scan(&expense.ID)
	return expense, err
}

func (er expenseRepository) GetExpenseById(id string) (models.Expense, error) {
	er.Ping()
	var expense models.Expense
	expense.ID = id
	err := er.db.QueryRow("SELECT category, price FROM expenses WHERE id=$1", id).Scan(&expense.Category, &expense.Price)
	return expense, err
}

func (er expenseRepository) DeleteExpenseById(id string) (string, error) {
	er.Ping()
	var foundId string
	err := er.db.QueryRow("DELETE FROM expenses WHERE id=$1 RETURNING id", id).Scan(&foundId)
	return foundId, err
}

func (er expenseRepository) Ping() {
	err := er.db.Ping()
	if err != nil {
		log.Printf("could not ping the db: %v\n", err)
	}
}
