package domain

import (
	"github.com/stretchr/testify/assert"
	"microservices-ex-app/mvc/services"
	"net/http"
	"testing"
)

func TestGetOrderNoOrderFound(t *testing.T) {
	// Initialization:

	// Execution:
	order, err := services.OrdersService.GetOrder(0)

	// Validation:
	//		Using stretchr/testify library
	assert.Nil(t, order, "we were not expecting a order with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not found", err.Code)
	assert.EqualValues(t, "Order 0 was not found", err.Message)

	//		Using go testing library
	if order != nil {
		t.Error("we were not expecting a order with id 0")
	}
	if err == nil {
		t.Error("we were expecting an error when order id is 0")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("we were expecting 404 when user is not found")
	}
}

func TestGetOrderNoError(t *testing.T) {
	order, err := services.OrdersService.GetOrder(123)

	assert.Nil(t, err)
	assert.NotNil(t, order)

	assert.EqualValues(t, 123, order.Id)
	assert.EqualValues(t, "Fede", order.PersonName)
	assert.EqualValues(t, "Melba", order.ProductName)
}
