package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rumiant348/finance-solver/controllers"
	"log"
)

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	e := controllers.NewExpenses()
	router.GET("/expenses", e.GetExpenses)
	router.GET("/expenses/:id", e.GetExpensesById)
	router.POST("/expenses", e.PostExpenses)
	router.DELETE("/expenses/:id", e.DeleteExpenseById)
	//router.OPTIONS("/expenses", e.Preflight)
	//router.OPTIONS("/expenses/:id", e.Preflight)

	a := controllers.NewAmounts()
	router.GET("/amounts", a.GetAmounts)
	router.GET("/amounts/:id", a.GetAmountsById)
	router.POST("/amounts", a.PostAmounts)
	router.DELETE("/amounts/:id", a.DeleteAmountById)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
