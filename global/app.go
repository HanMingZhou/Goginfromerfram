package global

import (
	"Goginfromerfram/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	Config      config.Configuration
	ConfigViper *viper.Viper
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
