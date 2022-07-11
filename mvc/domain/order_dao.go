package domain

import (
	"fmt"
	"microservices-ex-app/mvc/utils"
	"net/http"
)

var (
	orders = map[int64]*Order{
		123: &Order{Id: 1, PersonName: "Fede", ProductName: "Melba"},
	}
)

func GetOrder(orderId int64) (*Order, *utils.ApplicationError) {
	if order := orders[orderId]; order != nil {
		return order, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("Order %v was not found", orderId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
