package ports

import (
	"go-storage-s3/entities/models"
)

type FileDatabaseRepository interface {
	Save(data *models.FileModel) error
	Find(conditions interface{}, limit int, offset int) ([]*models.FileModel, error)
}
