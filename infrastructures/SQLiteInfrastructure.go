package infrastructures

import (
	"AmigoAutonomo/config"
	"AmigoAutonomo/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetLogDatabase() (*gorm.DB, error) {
	dbPath := config.GetPropFromServiceSection().LogDatabasePath
	db, err := gorm.Open(sqlite.Open(dbPath+"logError.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entities.LogError{})
	if err != nil {
		return nil, err
	}

	return db,err
}
