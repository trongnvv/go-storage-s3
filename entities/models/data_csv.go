package models

import (
	"gorm.io/gorm"
)

type DataCsvModel struct {
	gorm.Model
	Description string `json:"description,omitempty"`
	Industry    string `json:"industry,omitempty"`
	Level       string `json:"level,omitempty"`
	Size        string `json:"size,omitempty"`
	LineCode    string `json:"line_code,omitempty"`
	Value       string `json:"value,omitempty"`
}

func (DataCsvModel) TableName() string {
	return "data_csv"
}
