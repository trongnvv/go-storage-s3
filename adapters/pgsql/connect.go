package pgsql

import (
	"fmt"
	"go-storage-s3/common/log"
	"go-storage-s3/configs"
	"go-storage-s3/entities/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func Connect(config *configs.Config) *gorm.DB {
	cf := config.Postgresql
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s", cf.Host,
		cf.Port, cf.User, cf.DbName, cf.SslMode, cf.Password)
	fmt.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(err, "failed to connect database, db name:[%s]", cf.DbName)
	} else {
		log.Infof("connect to db :[%s]", cf.DbName)
	}
	psql, err := db.DB()
	if err != nil {
		log.Fatalf(err, "failed to connect database, db name:[%s]", cf.DbName)
		return nil
	}
	psql.SetConnMaxLifetime(time.Duration(cf.MaxLifeTime) * time.Second)
	if cf.AutoMigrate {
		if err = db.AutoMigrate(&models.FileModel{}); err != nil {
			log.Fatal(err, "auto migrate fail")
			return nil
		}
	}
	return db
}
