package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	healthResponse, err := http.Get("http://localhost:3000/api/health")
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	body, err := io.ReadAll(healthResponse.Body)
	if err != nil {
		t.Error("An error has ocurred", err)
	}
	var response string
	if err = json.Unmarshal(body, &response); err != nil {
		t.Error("An error has ocurred", err)
	}

	if response != "API Health" {
		t.Log("Body response not expected")
	}
}

func TestCreateAccount(t *testing.T) {
	bodyRequest, _ := json.Marshal(map[string]string{
		"document":       "12345678002",
		"name":           "user_test",
		"address":        "Rua 1",
		"address_number": "444",
		"uf":             "RS",
		"city":           "Porto alegre",
	})

	bodyRequestBytes := bytes.NewBuffer(bodyRequest)

	accountResponse, err := http.Post("http://localhost:3000/api/account", "application/json", bodyRequestBytes)
	if err != nil {
		t.Error("An error has ocurred", err)
	}

	body, err := io.ReadAll(accountResponse.Body)
	if err != nil {
		t.Error("An error has ocurred", err)
	}
	var response map[string]interface{}
	if err = json.Unmarshal(body, &response); err != nil {
		t.Error("An error has ocurred", err)
	}

	if accountResponse.StatusCode != 201 {
		t.Log("Http Status code not expected")
	}

	if response["id"].(float64) == 0 {
		t.Log("Body response not expected")
	}
}
