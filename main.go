package main

import (
	"urlshortener/config"
	"urlshortener/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8080"))
}
