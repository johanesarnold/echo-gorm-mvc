package models

import (
	"time"
)

type MasterParameter struct {
	ParameterID          string     `gorm:"primary_key;column:PARAMETER_ID;" json:"PARAMETER_ID"`
	ParameterLabel       string     `gorm:"column:PARAMETER_LABEL;" json:"PARAMETER_LABEL"`
	ParameterDescription string     `gorm:"column:PARAMETER_DESCRIPTION;" json:"PARAMETER_DESCRIPTION"`
	CreatedBy            *string    `gorm:"column:CREATED_BY" json:"CREATED_BY"`
	CreatedAt            *time.Time `gorm:"column:CREATED_AT" json:"CREATED_AT"`
	UpdatedBy            *string    `gorm:"column:UPDATED_BY" json:"UPDATED_BY"`
	UpdatedAt            *time.Time `gorm:"column:UPDATED_AT" json:"UPDATED_AT"`
	DeletedBy            *string    `gorm:"column:DELETE_BY" json:"DELETE_BY"`
	DeletedAt            *time.Time `gorm:"column:DELETE_AT" json:"DELETE_AT"`
}

func (MasterParameter) TableName() string {
	return "tm_master_parameter"
}
