package app

import (
	"github.com/spf13/viper"
	"github.com/uzzalhcse/GoSpire/app/exceptions"
	"github.com/uzzalhcse/GoSpire/config"
)

func Boot(v *viper.Viper) {
	config.Boot(v)
	err := BootDatabase()
	exceptions.PanicIfNeeded(err)
}
