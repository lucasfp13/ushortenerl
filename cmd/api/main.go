package main

import (
	"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/lucasfp13/ushortenerl/db"
	"github.com/lucasfp13/ushortenerl/handler"
)

func main() {
	godotenv.Load()
	err := db.MongoConnect()
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	defer func() {
		if err := db.Client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/r/", handler.RedirectHandler)

	log.Println("Service running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
