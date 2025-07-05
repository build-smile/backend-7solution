package infrastructure

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	CFG = Config{}
)

func InitConfig() {
	viper.SetConfigName("config")    // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&CFG)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	CFG.Jwt.SecretKey = []byte(CFG.Jwt.SecretKeyString)
}

type Config struct {
	App     App     `mapstructure:"app"`
	MongoDB Mongodb `mapstructure:"mongodb"`
	Jwt     Jwt     `mapstructure:"jwt"`
}
type App struct {
	Port int    `mapstructure:"port"`
	Name string `mapstructure:"name"`
}

type Mongodb struct {
	Uri                 string `mapstructure:"uri"`
	Database            string `mapstructure:"database"`
	ConnectTimeoutMilli int    `mapstructure:"connectTimeoutMilli"`
	ExecuteTimeoutMilli int    `mapstructure:"executeTimeoutMilli"`
}
type Jwt struct {
	SecretKeyString string `mapstructure:"secret-key"`
	SecretKey       []byte
}
