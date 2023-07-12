package common

import (
	"context"
	"encoding/json"
	"net/http"
)

func ResponseError(w http.ResponseWriter, ctx context.Context, key string, err error) {
	response := map[string]string{"msg": key}
	if err != nil {
		response = map[string]string{"msg": key, "error": err.Error()}
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(response)
}

func ResponseCreated(w http.ResponseWriter, response interface{}) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func ResponseSuccess(w http.ResponseWriter, response interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ResponseSimpleSuccess(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"success": true})
}

const (
	ErrorDatabasePersistence     string = "ErrorDatabasePersistence"
	ErrorDatabaseQuery           string = "ErrorDatabaseQuery"
	ErrorUnexpected              string = "ErrorUnexpected"
	JsonCantGetFromBody          string = "JsonCantGetFromBody"
	JsonNotDeserialized          string = "JsonNotDeserialized"
	IdNotReceivedFromQuerystring string = "IdNotReceivedFromQuerystring"
	BodyNotValidated             string = "BodyNotValidated"
)
