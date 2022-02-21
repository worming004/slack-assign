package assigner

import (
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AssignHistory struct {
	gorm.Model
	UserId       string
	DateAssigned time.Time
}

type PonderedAssigner struct {
	*gorm.DB
	excludedUserIds []string
	subAssign       Assigner
}

func NewPonderedAssigner(dbFolder string, usersToRemove []string) (PonderedAssigner, error) {
	dbPath := filepath.Join(dbFolder, "slack-history.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return PonderedAssigner{}, err
	}

	db.AutoMigrate(&AssignHistory{})

	return PonderedAssigner{db, usersToRemove, NewSimplerAssigner(usersToRemove)}, nil
}

func (pa PonderedAssigner) Assign(users []string) string {
	selectedUserId := pa.subAssign.Assign(users)
	assignHistoryToStore := AssignHistory{UserId: selectedUserId, DateAssigned: time.Now()}

	pa.Create(&assignHistoryToStore)

	return selectedUserId
}
