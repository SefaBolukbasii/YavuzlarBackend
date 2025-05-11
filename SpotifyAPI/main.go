package main

import (
	db "spotifyAPI/config"
	"spotifyAPI/routes"
)

func main() {
	db.Conf()
	routes.Routes()

}
