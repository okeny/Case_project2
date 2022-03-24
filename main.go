package main

import (
	"log"
	router "project-cases/router"
)

func main() {

	var address = ":3001"

	r := router.SetupRouter()
	log.Fatal(r.Run(address))
}
