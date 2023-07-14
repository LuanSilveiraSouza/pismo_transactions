package account

import "time"

type Account struct {
	Id            uint      `gorm:"primaryKey" json:"id,omitempty"`
	Document      string    `gorm:"not null" json:"document,omitempty"`
	Name          string    `gorm:"not null" json:"name,omitempty"`
	Balance       float64   `gorm:"not null" json:"balance,omitempty"`
	Address       string    `gorm:"not null" json:"address"`
	AddressNumber string    `gorm:"nut null" json:"address_number"`
	City          string    `gorm:"nut null" json:"city"`
	UF            string    `gorm:"nut null" json:"uf"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}
