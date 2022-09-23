package repositories

import (
	"go-storage-s3/core/ports"
	"go-storage-s3/entities/models"
	"gorm.io/gorm"
)

type FileCsvRepository struct {
	db *gorm.DB
}

func NewFileCsvRepository(db *gorm.DB) ports.FileCsvDatabaseRepository {
	return &FileCsvRepository{db}
}

func (r *FileCsvRepository) Save(data *models.DataCsvModel) error {
	tx := r.db.Model(data).Save(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
