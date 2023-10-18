package config

import (
	"github.com/spf13/viper"
	"github.com/uzzalhcse/GoSpire/app/exceptions"
)

func Boot(v *viper.Viper) {
	initAppConfig(v)
	initDBConfig(v)
	initJwt(v)
}
func InitConfig(fileName string) (*viper.Viper, error) {
	v := viper.New()
	v.AddConfigPath("./")
	v.SetConfigType("env")
	v.SetConfigName(fileName)
	v.AllowEmptyEnv(true)
	v.AutomaticEnv()

	err := v.ReadInConfig()
	exceptions.PanicIfNeeded(err)

	return v, nil
}
