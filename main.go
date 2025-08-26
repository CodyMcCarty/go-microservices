package main

import (
	"fmt"
	"log"

	"github.com/CodyMcCarty/go-microservices/internal/database"
	"github.com/CodyMcCarty/go-microservices/internal/server"
)

func main() {
	fmt.Println("Starting")

	db, err := database.NewDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to initialize database client: %s", err)
	}
	srv := server.NewEchoServer(db)
	if err := srv.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
