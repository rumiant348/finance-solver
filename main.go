package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rumiant348/finance-solver/controllers"
)

func main() {
	router := gin.Default()
	e := controllers.NewExpenses()
	router.GET("/expenses", e.GetExpenses)
	router.POST("/expenses", e.PostExpenses)
	router.GET("/expenses/:id", e.GetExpensesById)
	router.DELETE("/expenses/:id", e.DeleteExpenseById)
	router.OPTIONS("/expenses", e.Preflight)
	router.OPTIONS("/expenses/:id", e.Preflight)

	router.Run(":8080")
}
