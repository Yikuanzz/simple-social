package data

import (
	"fmt"

	"github.com/yikuanzz/social/internal/base/conf"
	"github.com/yikuanzz/social/internal/base/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Data struct {
	DB *gorm.DB
}

func NewData(db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
		sqlDB, _ := db.DB()
		if err := sqlDB.Close(); err != nil {
			log.Errorf("failed to close sqlDB, err: %s", err)
		}
	}
	return &Data{
		DB: db,
	}, cleanup, nil
}

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.DBConfigs.User, conf.DBConfigs.Password, conf.DBConfigs.Address, conf.DBConfigs.DBName)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
