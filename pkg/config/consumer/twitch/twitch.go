package twitch


type TwitchEventConfig struct {
	Name        string `yaml:"name"`
	EventSecret string `yaml:"eventSecret"`
	Url         string `yaml:"url"`
}
