package main

import (
	db "API/config"
	"API/routes"
	"net/http"
)

func main() {
	db.Conf()
	routes.Routes()
	http.ListenAndServe(":8080", nil)
}
