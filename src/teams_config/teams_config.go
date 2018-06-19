package teams_config

type TeamsConfig struct {
	SparkSecret string
	AccessToken string
	Username string
	BotId string
	TargetURL string
}

func NewConfig(conf map[string]string ) TeamsConfig {
	config := TeamsConfig{}
	config.SparkSecret = conf["SparkSecret"]
	config.AccessToken = conf["AccessToken"]
	config.Username = conf["Username"]
	config.BotId = conf["BotId"]
	return config
}