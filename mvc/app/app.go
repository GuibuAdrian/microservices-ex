package app

import (
	"microservices-ex-app/mvc/controller"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/orders", controller.GetOrder)

	if err := http.ListenAndServe("localhost:8084", nil); err != nil {
		panic(err)
	}

}
