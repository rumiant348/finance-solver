package repository

import (
	"database/sql"
	"github.com/rumiant348/finance-solver/models"
	"log"
)

type AmountRepository interface {
	CreateTable() error
	GetAmounts() ([]models.Amount, error)
	GetAmountById(id string) (models.Amount, error)
	CreateAmount(expense models.Amount) (models.Amount, error)
	DeleteAmountById(id string) (string, error)
}

type amountRepository struct {
	db *sql.DB
}

func NewAmountRepository(db *sql.DB) AmountRepository {
	return amountRepository{db: db}
}

func (r amountRepository) CreateTable() error {
	r.Ping()
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS amounts(
			id serial PRIMARY KEY, 
			price float NOT NULL,
    		"from" VARCHAR (3) NOT NULL, 
    		"to" VARCHAR (3) NOT NULL
	  );`)
	return err
}

func (r amountRepository) GetAmounts() ([]models.Amount, error) {
	r.Ping()
	rows, err := r.db.Query("SELECT * FROM amounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var amounts []models.Amount
	for rows.Next() {
		var a models.Amount
		if err := rows.Scan(&a.ID, &a.Price, &a.From, &a.To); err != nil {
			return amounts, err
		}
		amounts = append(amounts, a)
	}
	return amounts, nil
}

func (r amountRepository) CreateAmount(amount models.Amount) (models.Amount, error) {
	r.Ping()
	err := r.db.QueryRow("INSERT INTO amounts(price, \"from\", \"to\") VALUES($1, $2, $3)  RETURNING id",
		amount.Price, amount.From, amount.To).Scan(&amount.ID)
	return amount, err
}

func (r amountRepository) GetAmountById(id string) (models.Amount, error) {
	r.Ping()
	var amount models.Amount
	amount.ID = id
	err := r.db.QueryRow("SELECT price, \"from\", \"to\" FROM amounts WHERE id=$1", id).Scan(
		&amount.Price, &amount.From, &amount.To)
	return amount, err
}

func (r amountRepository) DeleteAmountById(id string) (string, error) {
	r.Ping()
	var foundId string
	err := r.db.QueryRow("DELETE FROM amounts WHERE id=$1 RETURNING id", id).Scan(&foundId)
	return foundId, err
}

func (r amountRepository) Ping() {
	err := r.db.Ping()
	if err != nil {
		log.Printf("could not ping the db: %v\n", err)
	}
}
