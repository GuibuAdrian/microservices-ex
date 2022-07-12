package controller

import (
	"github.com/gin-gonic/gin"
	"microservices-ex-app/mvcGinGonic/services"
	"microservices-ex-app/mvcGinGonic/utils"
	"net/http"
	"strconv"
)

func GetOrder(c *gin.Context) {
	orderId, err := strconv.ParseInt(c.Param("order_id"), 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "order_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, appErr)
		return
	}

	order, apiErr := services.OrdersService.GetOrder(orderId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, order)
}
