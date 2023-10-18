package config

import "github.com/spf13/viper"

var Jwt JwtConfig

type JwtConfig struct {
	SecretKey []byte
}

func initJwt(v *viper.Viper) {
	Jwt.SecretKey = []byte(v.GetString("JWT_SECRET_KEY"))
}
