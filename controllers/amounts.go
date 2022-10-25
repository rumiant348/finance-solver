package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/rumiant348/finance-solver/database"
	"github.com/rumiant348/finance-solver/models"
	"github.com/rumiant348/finance-solver/repository"
	"log"
	"net/http"
)

type Amounts struct {
	r  repository.AmountRepository
	db *sql.DB
}

func NewAmounts() *Amounts {
	db := database.Connection()
	r := repository.NewAmountRepository(db)
	err := r.CreateTable()
	if err != nil {
		log.Fatal(err)
	}
	return &Amounts{
		db: db,
		r:  r,
	}
}

func (a *Amounts) GetAmounts(c *gin.Context) {
	amounts, err := a.r.GetAmounts()
	if err != nil {
		log.Printf("Error while getting amounts: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error: " + err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, amounts)
}

func (a *Amounts) PostAmounts(c *gin.Context) {
	var newAmount models.Amount
	if err := c.BindJSON(&newAmount); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}
	amount, err := a.r.CreateAmount(newAmount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Internal error: "+err.Error())
		return
	}
	c.IndentedJSON(http.StatusCreated, amount)
}

func (a *Amounts) DeleteAmountById(c *gin.Context) {
	_, err := a.r.DeleteAmountById(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "amount with id " + c.Param("id") + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": c.Param("id")})
}

func (a *Amounts) GetAmountsById(c *gin.Context) {
	amount, err := a.r.GetAmountById(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "amount with id " + c.Param("id") + " not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, amount)
}
