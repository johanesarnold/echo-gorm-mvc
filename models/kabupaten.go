package models

import "time"

type Kabupaten struct {
	KabupatenID    string     `gorm:"primary_key;column:KABUPATEN_ID" json:"KABUPATEN_ID"`
	KabupatenLabel string     `gorm:"column:KABUPATEN_LABEL" json:"KABUPATEN_LABEL"`
	ProvinsiID     string     `gorm:"column:PROVINSI_ID" json:"PROVINSI_ID"`
	CreatedBy      *string    `gorm:"column:CREATED_BY" json:"CREATED_BY"`
	CreatedAt      *time.Time `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	UpdatedBy      *string    `gorm:"column:UPDATED_BY" json:"UPDATED_BY"`
	UpdatedAt      *time.Time `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
	DeletedBy      *string    `gorm:"column:DELETE_BY" json:"DELETE_BY"`
	DeletedAt      *time.Time `gorm:"column:DELETE_AT" json:"DELETE_AT"`
}

func (Kabupaten) TableName() string {
	return "tm_kabupaten"
}
