package handler

import (
	"fmt"
	"net/http"
)

type Order struct{}

func (order *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO: create an order")
}

func (order *Order) List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO: list all orders")
}

func (order *Order) GetById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO: get an order by id")
}

func (order *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO: update an order by id")
}

func (order *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TODO: delete an order by id")
}
