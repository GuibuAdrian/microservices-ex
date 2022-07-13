package app

import (
	"microservices-ex-app/mvcGinGonic/controller"
)

func mapUrls() {
	router.GET("/orders/:order_id", controller.GetOrder) //example: .../orders/123
}
