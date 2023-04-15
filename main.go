package main

import (
	"github.com/lemkova/mygram-h8/database"
	"github.com/lemkova/mygram-h8/router"
)

func main() {
	database.Connect()
	r := router.StartApp()
	r.Run(":8080")
}
