package main

import (
	"log"
	"mail-app/db"
	"mail-app/router"
)

func main() {
	r := router.Default()

	db.Init()

	log.Fatal(r.Run())
}
