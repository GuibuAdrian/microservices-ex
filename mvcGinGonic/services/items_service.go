package services

import (
	"microservices-ex-app/mvcGinGonic/domain"
	"microservices-ex-app/mvcGinGonic/utils"
	"net/http"
)

type itemService struct{}

var ItemService itemService

func (is *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
