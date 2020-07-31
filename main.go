package main

import (
	"github.com/golang-api/methods"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	log.Fatal(http.ListenAndServe(":8000", methods.Config()))
}
