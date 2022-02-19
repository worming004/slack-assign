package assigner

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type AssignHistory struct {
	gorm.Model
	UserId       string
	DateAssigned time.Time
}

type PonderedAssigner struct {
	*gorm.DB
}

func NewPonderedAssigner() (PonderedAssigner, error) {
	db, err := gorm.Open(sqlite.Open("slack-history.db"), &gorm.Config{})
	if err != nil {
		return PonderedAssigner{}, err
	}

	db.AutoMigrate(&AssignHistory{})

	return PonderedAssigner{db}, nil
}

func (pa PonderedAssigner) Assign(users []string) string {
	// TODO make history useful
	return ""
}
