package router

import (
	"net/http"
	"go-web/handler"
)

func Route() {
	http.HandleFunc("/", handler.HandleMainPage)
}