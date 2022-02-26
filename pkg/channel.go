package assign

import (
	"fmt"

	"github.com/slack-go/slack"
	"github.com/worming004/slack-assign/pkg/assigner"
)

type Assign struct {
	Configuration Configuration
	Client        *slack.Client
	assigner      assigner.Assigner
}

func NewAssign(C Configuration) (*Assign, error) {
	api := slack.New(C.Token)
	assign, err := assigner.NewPonderedAssigner(C.DbPath)
	if err != nil {
		return nil, err
	}
	return &Assign{C, api, assign}, nil
}

func (a *Assign) Run() error {
	users, err := a.GetUsers()
	if err != nil {
		return err
	}

	// the assign bot is registered in the channel. So we remove it
	users = withoutExcluded(users, []string{a.Configuration.AssignUserId})

	selectedUser := a.assigner.Assign(users)

	if !a.Configuration.IsDebug {
		a.PostMessage(selectedUser)
	} else {
		fmt.Println(selectedUser)
	}

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

func withoutExcluded(users []string, usersToRemove []string) []string {
	var withoutExcluded []string

	excluded := make(map[string]struct{})
	for _, u := range usersToRemove {
		excluded[u] = struct{}{}
	}

	for _, i := range users {
		if _, ok := excluded[i]; !ok {
			withoutExcluded = append(withoutExcluded, i)
		}
	}
	return withoutExcluded
}
