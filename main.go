package main

import (
	"fmt"
	"log"

	"github.com/CodyMcCarty/go-microservices/internal/database"
	"github.com/CodyMcCarty/go-microservices/internal/server"
)

/* additional criteria for future:
multiple dbs
main app with 1 or 2 microservices.
config file, along with other comments in *customer
*/

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
