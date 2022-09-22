package repositories

import (
	"errors"
	"go-storage-s3/core/ports"
	"go-storage-s3/entities/models"
	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) ports.FileRepository {
	return &FileRepository{db}
}

func (r *FileRepository) Save(data *models.FileModel) error {
	tx := r.db.Model(data).Save(data)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *FileRepository) Find(conditions interface{}, limit int, offset int) ([]*models.FileModel, error) {
	var records []*models.FileModel
	tx := r.db.Model(records).Order("id desc").Offset(offset).Limit(limit).Find(records, conditions)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, tx.Error
	}

	return records, nil
}
