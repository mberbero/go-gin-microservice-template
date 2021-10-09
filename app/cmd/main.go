package main

import (
	"os"

	"github.com/mberbero/go-microservice-template/pkg/infra/database"
	"github.com/mberbero/go-microservice-template/pkg/infra/server"
)

var (
	host   string = os.Getenv("HOST")
	dbname string = os.Getenv("DBNAME")
	port   string = os.Getenv("PORT")
)

func init() {
	if os.Getenv("PORT") == "" {
		port = "8080"
	}
	if os.Getenv("DBNAME") == "" {
		dbname = "go-microservice-template"
	}
	if os.Getenv("HOST") == "" {
		host = "127.0.0.1"
	}
	database.Connect(dbname)
}

func main() {
	server.Run(host, port)
}
