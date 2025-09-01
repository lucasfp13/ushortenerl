package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/lucasfp13/ushortenerl/db"
	"github.com/lucasfp13/ushortenerl/services"
	"go.mongodb.org/mongo-driver/bson"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[len("/r/"):]
	if shortUrl == "" {
		http.Error(w, "Missing URL!", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	var url services.URL
	err := db.URLCollection.FindOne(ctx, bson.M{"short": shortUrl}).Decode(&url)
	if err != nil {
		log.Println("Erro FindOne:", err)
		http.NotFound(w, r)
		return
	}

	_, _ = db.URLCollection.UpdateOne(ctx,
		bson.M{"short": shortUrl},
		bson.M{"$inc": bson.M{"clicks": 1}},
	)

	http.Redirect(w, r, url.Original, http.StatusFound)
}
