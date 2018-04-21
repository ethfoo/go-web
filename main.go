package main

import (
	"log"
	"net/http"
	"go-web/router"
)

func main() {
	router.Route()
	err := http.ListenAndServe(":8688", nil)
	if err!=nil {
		log.Fatalln("listen and server error: ", err)
	}
}