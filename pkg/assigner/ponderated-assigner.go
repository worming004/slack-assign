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
	subAssign Assigner
}

func NewPonderedAssigner(dbFolder string) (PonderedAssigner, error) {
	dbPath := filepath.Join(dbFolder, "slack-history.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return PonderedAssigner{}, err
	}

	db.AutoMigrate(&AssignHistory{})

	return PonderedAssigner{db, NewSimplerAssigner()}, nil
}

func (pa PonderedAssigner) Assign(users []string) string {
	var last20Assigned []AssignHistory
	multiplier := len(users)
	pa.Order("date_assigned desc").Limit(multiplier).Find(&last20Assigned)

	weightedAssigner := assignWithWeighting{last20Assigned, users, multiplier + 1, pa.subAssign}
	selectedUserId := weightedAssigner.Assign(users)

	assignHistoryToStore := AssignHistory{UserId: selectedUserId, DateAssigned: time.Now()}

	pa.Create(&assignHistoryToStore)

	return selectedUserId
}

type assignWithWeighting struct {
	assignHistories []AssignHistory
	users           []string
	multiplier      int
	subAssign       Assigner
}

func (aw assignWithWeighting) Assign(users []string) string {
	var fullList []string
	apparitions := aw.apparitionPerUser(aw.assignHistories)

	for i := 0; i < aw.multiplier; i++ {
		for _, user := range aw.users {
			if apparitions[user] <= i {
				fullList = append(fullList, user)
			}
		}
	}

	if len(fullList) == 0 {
		return aw.subAssign.Assign(users)
	}
	return aw.subAssign.Assign(fullList)
}

func (aw assignWithWeighting) apparitionPerUser(history []AssignHistory) map[string]int {
	result := make(map[string]int)
	for _, h := range history {
		result[h.UserId]++
	}

	return result
}
