package controller

import (
	"encoding/json"
	"microservices-ex-app/mvc/services"
	"microservices-ex-app/mvc/utils"
	"net/http"
	"strconv"
)

func GetOrder(resp http.ResponseWriter, req *http.Request) {
	orderIdParam := req.URL.Query().Get("order_id")
	orderId, err := strconv.ParseInt(orderIdParam, 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "order_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		jsonValue, _ := json.Marshal(appErr)
		resp.WriteHeader(appErr.StatusCode)
		resp.Write(jsonValue)
		//Just return the bad request to the client
		return
	}

	order, apiErr := services.GetOrder(orderId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		//Handle the error and return to the client
		return
	}

	//return order to client
	jsonValue, _ := json.Marshal(order)
	resp.Write(jsonValue)
}
