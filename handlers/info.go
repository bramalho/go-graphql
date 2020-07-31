package handlers

import (
	u "go-graphql/utils"
	"net/http"
)

// Info about the app
var Info = func(w http.ResponseWriter, r *http.Request) {
	u.Respond(w, map[string]interface{}{
		"status":  true,
		"message": "It works!",
	})
}
