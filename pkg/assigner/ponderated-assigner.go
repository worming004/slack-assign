package assigner

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const multiplier int = 5

type AssignHistory struct {
	gorm.Model
	UserId       string
	DateAssigned time.Time
}

type PonderedAssigner struct {
	*gorm.DB
	subAssign       Assigner
	userIdsToIgnore []string
}

func (pa *PonderedAssigner) AddUserIdToIgnore(userid string) {
	pa.userIdsToIgnore = append(pa.userIdsToIgnore, userid)
}

func NewPonderedAssigner(dbFolder string) (PonderedAssigner, error) {
	dbPath := filepath.Join(dbFolder, "slack-history.db")
	l := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{})
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{Logger: l})
	if err != nil {
		return PonderedAssigner{}, err
	}

	db.AutoMigrate(&AssignHistory{})

	return PonderedAssigner{db, NewSimplerAssigner(), []string{}}, nil
}

func (pa PonderedAssigner) Assign(users []string) string {
	var last20Assigned []AssignHistory
	lenUsers := len(users)
	pa.Order("date_assigned desc").Not("user_id IN ?", pa.userIdsToIgnore).Limit(lenUsers * multiplier).Find(&last20Assigned)

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
