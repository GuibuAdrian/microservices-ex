package services

import (
	"microservices-ex-app/mvc/domain"
	"microservices-ex-app/mvc/utils"
)

func GetOrder(orderId int64) (*domain.Order, *utils.ApplicationError) {
	return domain.GetOrder(orderId)
}
