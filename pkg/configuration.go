package assign

import (
	"os"
	"strconv"
)

type Configuration struct {
	Token        string
	ChannelId    string
	AssignUserId string
	DbPath       string
	IsDebug      bool
}

func GetConfigurationByEnvironmentVariable() Configuration {
	token := os.Getenv("SLACK_TOKEN")
	channelId := os.Getenv("SLACK_CHANNEL_ID")
	assignUserId := os.Getenv("SLACK_ASSIGN_USER_ID")
	dbPath := os.Getenv("SLACK_DB_PATH")
	isDebugStr := os.Getenv("SLACK_IS_DEBUG")

	isDebug, err := strconv.ParseBool(isDebugStr)
	if err != nil {
		isDebug = false
	}

	return Configuration{
		Token:        token,
		ChannelId:    channelId,
		AssignUserId: assignUserId,
		DbPath:       dbPath,
		IsDebug:      isDebug,
	}
}
