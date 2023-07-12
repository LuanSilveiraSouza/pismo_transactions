package transaction

import (
	"context"
	"encoding/json"
	"io"
	"math"
	"net/http"

	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/account"
	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/common"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Service interface {
	Add(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s service) List(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	params := mux.Vars(r)
	id := params["id"]

	var account account.Account
	result := s.db.First(&account, id)

	if result.Error != nil {
		common.ResponseError(w, ctx, common.ErrorDatabaseQuery, result.Error)
		return
	}

	var transactions []Transaction
	result = s.db.Where("account_id = ?", id).Find(&transactions)
	if result.Error != nil {
		common.ResponseError(w, ctx, common.ErrorDatabaseQuery, result.Error)
		return
	}

	common.ResponseSuccess(w, transactions)
}

func (s service) Add(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var transaction Transaction
	reqBody, err := io.ReadAll(r.Body)
	if err != nil || !json.Valid(reqBody) {
		common.ResponseError(w, ctx, common.JsonCantGetFromBody, err)
		return
	}
	if err = json.Unmarshal(reqBody, &transaction); err != nil {
		common.ResponseError(w, ctx, common.JsonNotDeserialized, err)
		return
	}

	var account account.Account
	result := s.db.First(&account, transaction.AccountId)

	if result.RowsAffected == 0 {
		common.ResponseError(w, ctx, "Account doesn't exists", result.Error)
		return
	}

	if transaction.TransactionTypeId != PAGAMENTO.Id {
		transaction.Amount = math.Copysign(transaction.Amount, -1)
	}

	tx := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		updatedBalance := account.Balance + transaction.Amount

		if err = tx.Model(&account).Update("balance", updatedBalance).Error; err != nil {
			return err
		}

		return nil
	})

	if tx != nil {
		common.ResponseError(w, ctx, common.ErrorDatabasePersistence, tx)
		return
	}

	common.ResponseCreated(w, transaction)
}
