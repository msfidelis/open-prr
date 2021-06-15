package orm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"open-prr/pkg/db"
	"open-prr/pkg/logger"
)

var orm *gorm.DB

func GetInstance() (*gorm.DB, error) {
	log := logger.Instance()
	if orm == nil {
		log.Info().
			Str("component", "database").
			Msg("Creating ORM Instance")
		sqlDB, err := db.GetConnection()
		if err != nil {
			return nil, err
		}
		orm, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}
	log.Debug().
		Str("component", "database").
		Msg("Retrieving ORM Singlethon")
	return orm, nil
}
