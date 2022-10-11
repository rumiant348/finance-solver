package repository

import (
	"database/sql"

	"github.com/rumiant348/finance-solver/models"
)

type ExpenseRepository interface {
	GetExpenses() ([]models.Expense , error)
    GetExpenseById(id int) (models.Expense, error)
    CreateExpense(expense models.Expense) (models.Expense , error)
}

type expenseRepository struct {
    db *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository {
    return expenseRepository{db: db}
}

func (er expenseRepository) GetExpenses() ([]models.Expense , error) {

	rows, err := er.db.Query("SELECT * FROM expense")
	if err != nil {
        return nil, err
    }
	defer rows.Close()
	
	var expenses []models.Expense
	for rows.Next() {
		var exp models.Expense
		if err := rows.Scan(&exp.ID, &exp.Amount, &exp.Description); err != nil {
			return expenses, err
		}
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

func (er expenseRepository) CreateExpense(expense models.Expense) (models.Expense, error) {
	err := er.db.QueryRow("INSERT INTO expense(amount, description) VALUES($1, $2)  RETURNING id", 
	    expense.Amount, expense.Description).Scan(&expense.ID)
	return expense, err
}

func (er expenseRepository) GetExpenseById(id int) (models.Expense, error) {
	var expense models.Expense
	expense.ID = id
	err := er.db.QueryRow("SELECT Amount, Description FROM expense WHERE id=$1", id).Scan(&expense.Amount, &expense.Description)
	return expense, err
}
