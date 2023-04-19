package main

import (
	"fmt"
	"os"

	"github.com/lemkova/mygram-h8/database"
	"github.com/lemkova/mygram-h8/router"
)

func main() {
	database.Connect()
	r := router.StartApp()
	var port string
	if os.Getenv("Env") == "Production" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	} else {
		port = ":8080"
	}
	r.Run(port)
}
