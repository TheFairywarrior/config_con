package environment

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Settings struct {
	ConfigFilePath string `env:"CONFIG_CONNECTORCONFIG_FILE_PATH"`
}

var settings Settings

func GetSettings() Settings {
	if settings == (Settings{}) {
		if err := env.Parse(&settings); err != nil {
			log.Fatal(err)
		}
	}
	return settings
}


