package models

import "time"

type Token struct {
	Email             string     `gorm:"primary_key;column:EMAIL"`
	Platform          string     `gorm:"column:PLATFORM"`
	Token             string     `gorm:"column:TOKEN"`
	NotificationToken string     `gorm:"column:NOTIFICATION_TOKEN"`
	CreatedBy         *string    `gorm:"column:CREATED_BY" json:"CREATED_BY"`
	CreatedAt         *time.Time `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	UpdatedBy         *string    `gorm:"column:UPDATED_BY" json:"UPDATED_BY"`
	UpdatedAt         *time.Time `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
	DeletedBy         *string    `gorm:"column:DELETE_BY" json:"DELETE_BY"`
	DeletedAt         *time.Time `gorm:"column:DELETE_AT" json:"DELETE_AT"`
	Formatted         string
}

func (Token) TableName() string {
	return "tm_token"
}
