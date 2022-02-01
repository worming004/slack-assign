package assign

import (
	"fmt"
	"github.com/slack-go/slack"
	"math/rand"
	"time"
)

type Assign struct {
	Configuration Configuration
	Client        *slack.Client
}

func NewAssign(C Configuration) *Assign {
	api := slack.New(C.Token)
	return &Assign{C, api}
}

func (a *Assign) Run() error {
	users, err := a.GetUsers()
	if err != nil {
		return err
	}

	selectedUser := string(GetRandomItem(users))

	a.PostMessage(selectedUser)

	return nil
}

func (a *Assign) PostMessage(userid string) error {
	msg := fmt.Sprintf("Hello <@%s>, your turn to post a question", userid)
	_, _, err := a.Client.PostMessage(a.Configuration.ChannelId, slack.MsgOptionText(msg, false))
	return err
}

func (a *Assign) GetUsers() ([]string, error) {

	users, _, err := a.Client.GetUsersInConversation(&slack.GetUsersInConversationParameters{
		ChannelID: a.Configuration.ChannelId,
	})
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetRandomItem(items []string) string {
	rand.Seed(time.Now().Unix())
	position := rand.Int() % len(items)

	return items[position]
}
