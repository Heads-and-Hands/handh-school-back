package handlers

import (
	"net/http"
	"encoding/json"
	"handh-school-back/database"
	"handh-school-back/models"
)

var GetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var payloadData = map[string]string{
		"message": "ok",
	}
	var payload, _ = json.Marshal(payloadData)
	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)
})

var PostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	var newRequest models.Request
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&newRequest)

	database.OrmProvider.CreateRequest(newRequest)

	var payloadData = map[string]string{
		"message": "ok",
	}
	var payload, _ = json.Marshal(payloadData)
	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)
})

