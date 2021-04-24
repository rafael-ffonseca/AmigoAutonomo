package entities

import "time"

type LogError struct {
	ID				uint		`gorm:"primarykey"`
	CreatedAt		time.Time	`gorm:"autoCreateTime"`
	Action		 	string
	Error			string
}
