package models

import (
	"gorm.io/gorm"
)

type Destination string
type Status string

const (
	S3Type Destination = "S3"
	Basic  Destination = "BASIC"
)

const (
	Init       Status = "INIT"
	Processing Status = "PROCESSING"
	Done       Status = "DONE"
	Error      Status = "ERROR"
)

type FileModel struct {
	gorm.Model
	Id        uint64 `json:"id,omitempty" gorm:"primaryKey"`
	Uuid      string `json:"uuid"`
	Path      string `json:"path"`
	Size      int64  `json:"size"`
	Type      string `json:"type"`
	CreatedBy string `json:"created_by"`

	Destination Destination `json:"destination"`
	Status      Status      `json:"status"`
}

func (FileModel) TableName() string {
	return "files"
}
