package database

import (
	"electronics-store-go/internal/app/models"
	"electronics-store-go/internal/config"
	"electronics-store-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
		log = logger.NewLogger()
	)
	db, err = gorm.Open(postgres.Open(cfg.Db.DataSourceName), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	rawDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	log.Info("Database connected!")
	rawDb.SetMaxOpenConns(cfg.Db.Pool.MaxOpen)
	rawDb.SetMaxIdleConns(cfg.Db.Pool.MaxIdle)
	rawDb.SetConnMaxIdleTime(cfg.Db.Pool.MaxLifeTime)

	if cfg.Db.Migrate.Enable {
		err := db.AutoMigrate(&models.Product{}, &models.Category{}, &models.Brand{})
		if err != nil {
			return nil, err
		}
	}

	return db, nil

}
