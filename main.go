package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rumiant348/finance-solver/database"
	"github.com/rumiant348/finance-solver/models"
	"github.com/rumiant348/finance-solver/repository"
)

var expenseRepository repository.ExpenseRepository 
var db *sql.DB

func init() {
	db = database.Connection()
	expenseRepository = repository.NewExpenseRepository(db)
}

func main() {
	router := gin.Default()
	
	router.GET("/expenses", getExpenses)
	router.POST("/expenses", postExpenses)
	router.GET("/expenses/:id", getExpensesById)

	router.Run("localhost:8080")
	defer db.Close()
}


func getExpenses(c *gin.Context) {

	expenses, err := expenseRepository.GetExpenses();
	if err != nil {
		log.Printf("Error while getting expenses: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
		return
	}


	c.IndentedJSON(http.StatusOK, expenses)
}

func postExpenses(c *gin.Context) {
	var newExpense models.Expense

	if err := c.BindJSON(&newExpense); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error")
		return
	}

	newExpense, err := expenseRepository.CreateExpense(newExpense)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error")
		return
	}

	c.IndentedJSON(http.StatusCreated, newExpense)
}

func getExpensesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id should be a integer "})
		return
	}
	expense, err := expenseRepository.GetExpenseById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "expense with id " + strconv.Itoa(id) + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, expense)
	
}
