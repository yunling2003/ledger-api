package main

import (
	"ledger-api/src/config"
	"ledger-api/src/router"

	"log"
	"os"
)

func main() {
	r := router.Route()

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = config.All["port"]
	}

	log.Println("API listening at http://localhost:" + port)

	r.Run(":" + port)
}
