package assign

import "os"

type Configuration struct {
	Token        string
	ChannelId    string
	AssignUserId string
}

func GetConfigurationByEnvironmentVariable() Configuration {
	token := os.Getenv("SLACK_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")
	assignUserId := os.Getenv("SLACK_ASSIGN_USER_ID")

	return Configuration{
		Token:        token,
		ChannelId:    channelId,
		AssignUserId: assignUserId,
	}
}
