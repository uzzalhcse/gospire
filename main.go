package main

import (
	"github.com/uzzalhcse/GoSpire/app/exceptions"
	"github.com/uzzalhcse/GoSpire/bootstrap/app"
	"github.com/uzzalhcse/GoSpire/bootstrap/server"
	"github.com/uzzalhcse/GoSpire/config"
)

func main() {
	viper, err := config.InitConfig(".env")
	exceptions.PanicIfNeeded(err)

	app.Boot(viper)
	server.RunServer()

}
