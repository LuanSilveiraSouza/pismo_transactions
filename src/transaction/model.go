package transaction

import (
	"time"

	"github.com/LuanSilveiraSouza/pismo_transactions.git/src/account"
	"gorm.io/gorm"
)

type Transaction struct {
	Id                uint            `gorm:"primaryKey" json:"id"`
	AccountId         uint            `gorm:"not null" json:"account_id"`
	Account           account.Account `json:"account,omitempty"`
	TransactionTypeId uint            `gorm:"not null" json:"transaction_type_id,omitempty"`
	Type              TransactionType `gorm:"foreignKey:TransactionTypeId;references:Id" json:"type,omitempty"`
	Amount            float64         `gorm:"not null" json:"amount"`
	CreatedAt         time.Time       `json:"created_at"`
}

type TransactionType struct {
	Id          uint   `gorm:"primaryKey" json:"id,omitempty"`
	Description string `gorm:"not null" json:"description,omitempty"`
}

func TransactionTypeSeed(db *gorm.DB) error {
	types := []*TransactionType{&COMPRA_A_VISTA, &COMPRA_PARCELADA, &SAQUE, &PAGAMENTO}

	result := db.Save(types)

	return result.Error
}

var (
	COMPRA_A_VISTA   = TransactionType{Id: 1, Description: "COMPRA_A_VISTA"}
	COMPRA_PARCELADA = TransactionType{Id: 2, Description: "COMPRA_PARCELADA"}
	SAQUE            = TransactionType{Id: 3, Description: "SAQUE"}
	PAGAMENTO        = TransactionType{Id: 4, Description: "PAGAMENTO"}
)
