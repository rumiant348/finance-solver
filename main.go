package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rumiant348/finance-solver/database"
	"github.com/rumiant348/finance-solver/models"
	"github.com/rumiant348/finance-solver/repository"
)

var expenseRepository repository.ExpenseRepository
var db *sql.DB

func getExpenses(c *gin.Context) {
	db = database.Connection()
	expenseRepository = repository.NewExpenseRepository(db)
	defer db.Close()

	c.Header("Access-Control-Allow-Origin", "*")
	expenses, err := expenseRepository.GetExpenses()
	if err != nil {
		log.Printf("Error while getting expenses: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error: " + err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, expenses)
}

func postExpenses(c *gin.Context) {
	db = database.Connection()
	expenseRepository = repository.NewExpenseRepository(db)
	defer db.Close()

	var newExpense models.Expense

	c.Header("Access-Control-Allow-Origin", "*")
	if err := c.BindJSON(&newExpense); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}

	newExpense, err := expenseRepository.CreateExpense(newExpense)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, newExpense)
}

func deleteExpenseById(c *gin.Context) {
	db = database.Connection()
	expenseRepository = repository.NewExpenseRepository(db)
	defer db.Close()

	c.Header("Access-Control-Allow-Origin", "*")

	_, err := expenseRepository.DeleteExpenseById(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense with id " + c.Param("id") + " not found"})
		// c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}
	// c.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense with id " + id + " not found"})
	c.IndentedJSON(http.StatusOK, "")
}

func getExpensesById(c *gin.Context) {
	db = database.Connection()
	expenseRepository = repository.NewExpenseRepository(db)
	defer db.Close()

	c.Header("Access-Control-Allow-Origin", "*")

	expense, err := expenseRepository.GetExpenseById(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense with id " + c.Param("id") + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, expense)

}

func preflight(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Status(http.StatusNoContent)
}

func main() {
	router := gin.Default()

	router.GET("/expenses", getExpenses)
	router.POST("/expenses", postExpenses)
	router.GET("/expenses/:id", getExpensesById)
	router.DELETE("/expenses/:id", deleteExpenseById)
	router.OPTIONS("/expenses", preflight)

	router.Run(":8080")
	defer db.Close()
}
