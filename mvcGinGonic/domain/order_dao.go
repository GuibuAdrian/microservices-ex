package domain

import (
	"fmt"
	"log"
	"microservices-ex-app/mvcGinGonic/utils"
	"net/http"
)

var (
	orders = map[int64]*Order{
		123: {Id: 123, PersonName: "Fede", ProductName: "Melba"},
	}

	OrderDao ordersDaoInterface
)

func init() {
	OrderDao = &orderDao{}
}

type ordersDaoInterface interface {
	GetOrder(orderId int64) (*Order, *utils.ApplicationError)
}

type orderDao struct{}

func (od *orderDao) GetOrder(orderId int64) (*Order, *utils.ApplicationError) {
	log.Println("we're accessing the database")
	if order := orders[orderId]; order != nil {
		return order, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("Order %v was not found", orderId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
