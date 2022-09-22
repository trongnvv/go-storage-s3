package health

import (
	"go-storage-s3/common/log"
	"gorm.io/gorm"
)

type IHealth interface {
	Check(str string) error
}

type HealthcheckService struct {
	db *gorm.DB
}

func NewHealthService(db *gorm.DB) *HealthcheckService {
	return &HealthcheckService{
		db,
	}
}

func (h *HealthcheckService) Check(serviceName string) error {
	log.Infof("HealthCheck.Check start check DB")

	err := h.CheckDB()

	if err != nil {
		log.Infof("HealthCheck.Check db.Ping err: %v", err)
		return err
	}

	if err != nil {
		log.Infof("HealthCheck.Check cache.Ping err: %v", err)
		return err
	}

	return nil
}

func (h *HealthcheckService) CheckDB() error {
	db, err := h.db.DB()

	if err != nil {
		log.Infof("HealthCheck.Check db err: %v", err)
		return err
	}

	return db.Ping()
}
