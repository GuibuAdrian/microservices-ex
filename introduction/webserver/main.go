package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))
	})

	if err := http.ListenAndServe("localhost:8084", nil); err != nil { //Only works in the scope of this condition
		panic(err)
	}
}
