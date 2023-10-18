package config

import "github.com/spf13/viper"

var App AppConfig

type AppConfig struct {
	Name  string
	Host  string
	Port  string
	Debug bool
}

func initAppConfig(v *viper.Viper) {
	App.Name = v.GetString("APP_NAME")
	App.Host = v.GetString("APP_HOST")
	App.Port = v.GetString("APP_PORT")
	App.Debug = v.GetBool("APP_DEBUG")

}
