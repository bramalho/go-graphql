package handlers

import (
	"context"
	"go-graphql/models"
	u "go-graphql/utils"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Info about the app
var Info = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, map[string]interface{}{
		"status":  true,
		"app":     os.Getenv("APP_NAME"),
		"version": os.Getenv("APP_VERSION"),
	})
}

// Health for the services
var Health = func(w http.ResponseWriter, r *http.Request) {
	ok := true

	db := models.GetDB()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	dberr := db.Client().Ping(ctx, readpref.Primary())

	if dberr != nil {
		ok = false
	}

	message := "OK"
	if ok == false {
		message = "ERROR"
	}

	u.Respond(w, map[string]interface{}{
		"status":  ok,
		"message": message,
		"services": map[string]interface{}{
			"db": dberr == nil,
		},
	})
}
