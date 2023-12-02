package global

import (
	"Goginfromerfram/config"

	"github.com/spf13/viper"
)

type Application struct {
	Config      config.Configuration
	ConfigViper *viper.Viper
}

var App = new(Application)
