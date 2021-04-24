package entities

import (
	"time"
)

type Users struct {
	ID        			uint `gorm:"primarykey"`
	CreatedAt			time.Time	`gorm:"autoCreateTime"`
	Username 			string
	Password 			string
	FirstName			string
	LastName			string
	Email				string
	Phones				[]UserPhones `gorm:"foreignKey:UserId"`
	Address				UserAddress `gorm:"foreignKey:UserId"`
	Document			string

}
