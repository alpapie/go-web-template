package routes

import (
	"goupie-tracker/handler"
	"net/http"
)

func Route() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/artist", handler.Index)
	http.HandleFunc("/event", handler.Index)
}
