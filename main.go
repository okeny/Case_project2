package main

import (
	"log"
	router "project-cases/routing"
)

func main() {

	var address = ":3001"
	r := router.SetupRouter()
	//_ = utility.GetConnection()
	log.Fatal(r.Run(address))
}
