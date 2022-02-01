package assign

import "os"

type Configuration struct {
	Token     string
	ChannelId string
}

func GetConfigurationByEnvironmentVariable() Configuration {
	token := os.Getenv("SLACK_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")

	return Configuration{
		Token:     token,
		ChannelId: channelId,
	}
}
