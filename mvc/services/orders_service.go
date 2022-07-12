package services

import (
	"microservices-ex-app/mvc/domain"
	"microservices-ex-app/mvc/utils"
)

type ordersService struct {
}

var OrdersService ordersService

func (ordersSer *ordersService) GetOrder(orderId int64) (*domain.Order, *utils.ApplicationError) {
	order, err := domain.OrderDao.GetOrder(orderId)
	if err != nil {
		return nil, err
	}

	return order, nil
}
