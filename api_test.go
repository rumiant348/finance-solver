package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/rumiant348/finance-solver/models"
)

// var getUrl = "https://finance-solver-api-v2.fly.dev/expenses"
var getUrl = "http://localhost:8080/expenses"

func TestGetAll(t *testing.T) {
	resp, err := http.Get(getUrl)
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer resp.Body.Close()

	var expenses []models.Expense
	err = json.NewDecoder(resp.Body).Decode(&expenses)
	if err != nil {
		t.Fatalf(err.Error())
	}

	fmt.Println(expenses)
}

type userDTO struct {
	Category string
	Price    float32
}

func TestPost(t *testing.T) {

	u := userDTO{
		Category: "Страховка",
		Price:    3.5,
	}

	result, err := postExpense(u)
	if err != nil {
		t.Fatalf("error sending post: " + err.Error())
	}
	fmt.Println(result)
	assert.NotEqual(t, result.ID, 0)
	assert.Equal(t, result.Category, "Страховка")
	assert.Equal(t, result.Price, float32(3.5))
}

func TestDelete(t *testing.T) {
	u := userDTO{
		Category: "тест delete",
		Price:    340.50,
	}

	ex, err := postExpense(u)
	if err != nil {
		t.Fatalf("failed to post: " + err.Error())
	}
	statusCode, err := deleteExpense(ex.ID)
	if err != nil {
		t.Fatalf("failed to delete: " + err.Error())
	}
	assert.Equal(t, statusCode, 200)

	statusCode, err = deleteExpense(ex.ID)
	if err != nil {
		t.Fatalf("failed to delete: " + err.Error())
	}
	assert.Equal(t, statusCode, 404)

}

func postExpense(u userDTO) (models.Expense, error) {
	var result models.Expense

	m, err := json.Marshal(&u)
	if err != nil {
		return result, err
	}

	resp, err := http.Post(getUrl, "application/json", bytes.NewBuffer(m))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

func deleteExpense(id string) (int, error) {
	req, err := http.NewRequest("DELETE", getUrl+"/"+id, nil)
	if err != nil {
		return 0, err
	}
	resp, err := http.DefaultClient.Do(req)
	return resp.StatusCode, err
}
