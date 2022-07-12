package services

import (
	"github.com/stretchr/testify/assert"
	"log"
	"microservices-ex-app/mvcGinGonic/domain"
	"microservices-ex-app/mvcGinGonic/utils"
	"net/http"
	"testing"
)

var (
	orderDaoMock     ordersDaoMock
	getOrderFunction func(orderId int64) (*domain.Order, *utils.ApplicationError)
)

func init() {
	domain.OrderDao = &ordersDaoMock{}
}

type ordersDaoMock struct{}

func (orM *ordersDaoMock) GetOrder(orderId int64) (*domain.Order, *utils.ApplicationError) {
	return getOrderFunction(orderId)
}

func TestGetOrderNotFoundInDataBase(t *testing.T) {
	getOrderFunction = func(orderId int64) (*domain.Order, *utils.ApplicationError) {
		log.Println("Mocked GetOrder")
		return nil, &utils.ApplicationError{
			Message:    "Order 0 was not found",
			StatusCode: http.StatusNotFound,
			Code:       "",
		}
	}
	order, err := OrdersService.GetOrder(0)

	assert.Nil(t, order)
	assert.NotNil(t, err)
	assert.EqualValues(t, "Order 0 was not found", err.Message)
}

func TestGetOrderNoError(t *testing.T) {
	getOrderFunction = func(orderId int64) (*domain.Order, *utils.ApplicationError) {
		log.Println("Mocked GetOrder")
		return &domain.Order{
			Id: 123,
		}, nil
	}
	order, err := OrdersService.GetOrder(123)

	assert.Nil(t, err)
	assert.NotNil(t, order)
	assert.EqualValues(t, 123, order.Id)
}
