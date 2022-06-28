package main

import (
	"electro3-project-go/db"
	"electro3-project-go/router"
	"log"
)

func main() {
	r := router.Default()

	db.Init()

	log.Fatal(r.Run(":8080"))
}
