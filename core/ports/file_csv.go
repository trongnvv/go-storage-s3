package ports

import (
	"go-storage-s3/entities/models"
)

type FileCsvDatabaseRepository interface {
	Save(data *models.DataCsvModel) error
}
