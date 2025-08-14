package config

import (
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

//reflection:understanding structure of data
type Config struct{
	Primary Primary `koanf:"primary" validate:"required"`
	Server ServerConfig `koanf:"server" validate:"required"`
	DataBase DataBaseConfig `koanf:"database" validate:"required"`
	Auth AuthConfig `koanf:"auth" validate:"required"`
	Redis RedisConfig `koanf:"redis" validate:"required"`
}

type Primary struct{
	Env string `koanf:"env" validate:"required"`
}

type ServerConfig struct{
	Port string `koanf:"port" validate:"required"`
	ReadTimeout string `koanf:"ReadTimeout" validate:"required"`
	WriteTimeout string `koanf:"WriteTimeout" validate:"required"`
	IdleTimeout string `koanf:"IdleTimeout" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate:"required"`
}

type DataBaseConfig struct{
	Host string `koanf:"host" validate:"required"`
	Port string  `koanf:"port" validate:"required"`
	User string	`koanf:"user" validate:"required"`
	Password string `koanf:"password"`
	Name string `koanf:"name" validate:"required"`
	SSLMode string `koanf:"ssl_mode" validate:"required"`
	MaxOpenConns int `koanf:"max_open_conns" validate:"required"`
	MaxIdleConns int `koanf:"max_idle_conns" validate:"required"`
	ConnMaxLifetime int `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdleTime int `koanf:"conn_max_idle_time" validate:"required"`
}



type AuthConfig struct{
	SecretKey string `koanf:"secret_key" validate:"required"`
}

type RedisConfig struct{
	Address string `koanf:"address" validate:"required"`
}


func LoadConfig()(*Config,error){
	logger:=zerolog.New(zerolog.ConsoleWriter{Out:os.Stderr}).With().Timestamp().Logger()

	k:=koanf.New(".")

	err:=k.Load(env.Provider("BIOLERPLATE_",".",func (s string) string{
		return strings.ToLower(strings.TrimPrefix(s,"BIOLERPLATE_"))
	}),nil)

	if(err!=nil){
		logger.Fatal().Err(err).Msg("couldnot load initial variables")
	}
	mainconfig:=&Config{}
	err=k.Unmarshal("",mainconfig)
	if(err!=nil){
		logger.Fatal().Err(err).Msg("couldnot unmarshall main config")
	}

	validate:=validator.New()
	err=validate.Struct(mainconfig)
	if(err!=nil){
		logger.Fatal().Err(err).Msg("config validation failed....")
	}

	return mainconfig,nil
}



