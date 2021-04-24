package infrastructures

import (
	"AmigoAutonomo/config"
	"AmigoAutonomo/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase(Homologation bool) (*gorm.DB, error) {
	dbUser := config.GetPropFromServiceSection().DatabaseUser
	dbPass := config.GetPropFromServiceSection().DatabasePassword
	dbHost := config.GetPropFromServiceSection().DatabaseHost
	dbPort := config.GetPropFromServiceSection().DatabasePort
	var dbName string
	if !Homologation {
		dbName = config.GetPropFromServiceSection().DatabaseName + "Prd"
	} else {
		dbName = config.GetPropFromServiceSection().DatabaseName + "Hmlg"
	}
	db, err := gorm.Open(
		mysql.New(mysql.Config{DSN: dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8&parseTime=True&loc=Local"}),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entities.Users{},
		&entities.UserAddress{},
		&entities.UserPhones{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}