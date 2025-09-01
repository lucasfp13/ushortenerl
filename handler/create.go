package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/lucasfp13/ushortenerl/db"
	"github.com/lucasfp13/ushortenerl/services"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only!", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Original string `json:"original"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid URL!", http.StatusBadRequest)
		return
	}

	if input.Original == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortUrl := generateCode(6)
	url := &services.URL{
		Original:  input.Original,
		Short:     shortUrl,
		CreatedAt: time.Now(),
		Clicks:    0,
	}

	_, err := db.URLCollection.InsertOne(context.Background(), url)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Create error!", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(url)
}

func generateCode(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
