package account

import "time"

type Account struct {
	Id        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Document  string    `gorm:"not null" json:"document,omitempty"`
	Name      string    `gorm:"not null" json:"name,omitempty"`
	Balance   float64   `gorm:"not null" json:"balance,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
