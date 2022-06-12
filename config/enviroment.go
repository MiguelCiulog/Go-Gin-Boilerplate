package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper
var Env *Enviroment

type Enviroment struct {
	DB struct {
		DatabaseName     string `mapstructure:"database"`
		DatabaseUrl      string `mapstructure:"url"`
		DatabaseUser     string `mapstructure:"user"`
		DatabasePassword string `mapstructure:"password"`
		DatabasePort     string `mapstructure:"port"`
	} `mapstructure:"db"`

	API struct {
		ApiPort string `mapstructure:"port"`
		ApiUrl  string `mapstructure:"url"`
	} `mapstructure:"api"`
}

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	viper.SetConfigName("development")
	viper.AddConfigPath("config/")
	err = viper.ReadInConfig()

	if err != nil {
		fmt.Printf("%v", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}
}

// Returns viper config
func GetConfig() *viper.Viper {
	return config
}
