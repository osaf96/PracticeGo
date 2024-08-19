package main

import (
	"log"

	"github.com/osaf96/PRACTICEGO/internal/server"
)

func main() {
	// Code
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
