package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

var getUrl = "https://finance-solver-api-v2.fly.dev/expenses"

func TestMain(t *testing.T) {
	resp, err := http.Get(getUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close() 

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}

