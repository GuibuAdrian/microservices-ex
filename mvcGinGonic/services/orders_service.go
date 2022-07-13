package services

import (
	"microservices-ex-app/mvcGinGonic/domain"
	"microservices-ex-app/mvcGinGonic/utils"
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
