package main

import (
	"fmt"
	"log"

	"github.com/baselrabia/distributed-services-with-go/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	fmt.Println("Starting New Http Server on: 0.0.0.0:8080")
	log.Fatal(srv.ListenAndServe())
}