package account

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/common"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Service interface {
	Add(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s service) Add(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var account Account
	reqBody, err := io.ReadAll(r.Body)
	if err != nil || !json.Valid(reqBody) {
		common.ResponseError(w, ctx, common.JsonCantGetFromBody, err)
		return
	}
	if err = json.Unmarshal(reqBody, &account); err != nil {
		common.ResponseError(w, ctx, common.JsonNotDeserialized, err)
		return
	}

	var accountExists Account
	result := s.db.Limit(1).Where("document = ?", account.Document).Find(&accountExists)

	if result.RowsAffected > 0 {
		common.ResponseError(w, ctx, "Record already exists", nil)
		return
	}

	if account.Name == "" {
		common.ResponseError(w, ctx, "Name is empty", nil)
		return
	}
	if account.Document == "" {
		common.ResponseError(w, ctx, "Document is empty", nil)
		return
	}
	if account.Address == "" {
		common.ResponseError(w, ctx, "Address is empty", nil)
		return
	}
	if account.AddressNumber == "" {
		common.ResponseError(w, ctx, "Address Number is empty", nil)
		return
	}
	if account.City == "" {
		common.ResponseError(w, ctx, "City is empty", nil)
		return
	}

	var UFTypes = []string{
		"AC",
		"AL",
		"AP",
		"AM",
		"BA",
		"CE",
		"DF",
		"ES",
		"GO",
		"MA",
		"MT",
		"MS",
		"MG",
		"PA",
		"PB",
		"PR",
		"PE",
		"PI",
		"RJ",
		"RN",
		"RS",
		"RO",
		"RR",
		"SC",
		"SP",
		"SE",
		"TO",
	}

	var validUf = false
	for _, uftype := range UFTypes {
		if uftype == account.UF {
			validUf = true
		}
	}

	if !validUf {
		common.ResponseError(w, ctx, "Invalid UF", nil)
		return
	}

	account.Balance = 0
	result = s.db.Create(&account)

	if result.Error != nil {
		common.ResponseError(w, ctx, common.ErrorDatabasePersistence, err)
		return
	}

	common.ResponseCreated(w, account)
}

func (s service) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	params := mux.Vars(r)
	id := params["id"]

	var account Account
	result := s.db.First(&account, id)

	if result.Error != nil {
		common.ResponseError(w, ctx, common.ErrorDatabaseQuery, result.Error)
		return
	}

	common.ResponseSuccess(w, account)
}
