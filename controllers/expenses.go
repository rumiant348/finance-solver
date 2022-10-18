package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rumiant348/finance-solver/database"
	"github.com/rumiant348/finance-solver/models"
	"github.com/rumiant348/finance-solver/repository"
)

type Expenses struct {
	expenseRepository repository.ExpenseRepository
	db                *sql.DB
}

func NewExpenses() *Expenses {
	db := database.Connection()
	expenseRepository := repository.NewExpenseRepository(db)
	return &Expenses{
		db:                db,
		expenseRepository: expenseRepository,
	}
}

func (e *Expenses) GetExpenses(c *gin.Context) {
	e.db.Ping()
	c.Header("Access-Control-Allow-Origin", "*")
	expenses, err := e.expenseRepository.GetExpenses()
	if err != nil {
		log.Printf("Error while getting expenses: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, expenses)
}

func (e *Expenses) PostExpenses(c *gin.Context) {
	e.db.Ping()
	c.Header("Access-Control-Allow-Origin", "*")

	var newExpense models.Expense

	if err := c.BindJSON(&newExpense); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}

	newExpense, err := e.expenseRepository.CreateExpense(newExpense)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, newExpense)
}

func (e *Expenses) DeleteExpenseById(c *gin.Context) {
	e.db.Ping()
	c.Header("Access-Control-Allow-Origin", "*")

	_, err := e.expenseRepository.DeleteExpenseById(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense with id " + c.Param("id") + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, "")
}

func (e *Expenses) GetExpensesById(c *gin.Context) {
	e.db.Ping()

	c.Header("Access-Control-Allow-Origin", "*")

	expense, err := e.expenseRepository.GetExpenseById(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense with id " + c.Param("id") + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, expense)

}

func (e *Expenses) Preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Status(http.StatusNoContent)
}
