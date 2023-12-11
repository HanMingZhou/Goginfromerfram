package global

import (
	"Goginfromerfram/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	Config      config.Configuration
	ConfigViper *viper.Viper
	Log         *zap.Logger
}

var App = new(Application)
