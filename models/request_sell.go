package models

import "time"

type RequestSell struct {
	RequestSellID string     `gorm:"primary_key;column:RS_ID" json:"RS_ID"`
	KabupatenID   string     `gorm:"column:KABUPATEN_ID" json:"KABUPATEN_ID"`
	Kabupaten     Kabupaten  `gorm:"foreignkey:KABUPATEN_ID;association_foreignkey:KABUPATEN_ID"`
	CreatedBy     *string    `gorm:"column:CREATED_BY" json:"CREATED_BY"`
	CreatedAt     *time.Time `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	UpdatedBy     *string    `gorm:"column:UPDATED_BY" json:"UPDATED_BY"`
	UpdatedAt     *time.Time `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
	DeletedBy     *string    `gorm:"column:DELETE_BY" json:"DELETE_BY"`
	DeletedAt     *time.Time `gorm:"column:DELETE_AT" json:"DELETE_AT"`
}

func (RequestSell) TableName() string {
	return "tr_jubel_request_sell"
}
