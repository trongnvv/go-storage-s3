package gracefully

import (
	"go-storage-s3/common/log"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type GracefulShutdownService struct {
	db *gorm.DB
}

func NewGracefulShutdownService(db *gorm.DB) *GracefulShutdownService {
	return &GracefulShutdownService{
		db,
	}
}

func (g *GracefulShutdownService) GracefulStop() {
	db, err := g.db.DB()

	if err == nil {
		if err := db.Close(); err != nil {
			log.Warnf("[GracefulShutdown] Close db connection err: %v", err)
		} else {
			log.Infof("[GracefulShutdown] DB connection gracefully closed")
		}
	} else {
		log.Infof("[GracefulShutdown] Get db err: %v", err)
	}
	return
}

func LoadDI() fx.Option {
	return fx.Provide(NewGracefulShutdownService)
}
