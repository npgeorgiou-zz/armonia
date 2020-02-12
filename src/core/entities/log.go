package entities

import (
	"time"
)

type Log struct {
	ID 		uint				`gorm:"primary_key"`
	Project string				`gorm:"size:10"`
	Level   string				`gorm:"size:10"`
	Message string				`gorm:"size:100"`
	SentAt time.Time
	CreatedAt time.Time
}
