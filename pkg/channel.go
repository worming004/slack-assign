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

func NewAssign(C Configuration) *Assign {
	api := slack.New(C.Token)
	assignUser := C.AssignUserId
	return &Assign{C, api, assigner.NewSimplerAssigner([]string{assignUser})}
}

func (a *Assign) Run() error {
	users, err := a.GetUsers()
	if err != nil {
		return err
	}

	selectedUser := a.assigner.Assign(users)

	//a.PostMessage(selectedUser)
	fmt.Println(selectedUser)
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
