package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var cfg *Config
var once sync.Once

func Init(path string) {
	once.Do(func() {
		conf, err := ReadConfig(path)
		if err != nil {
			panic(err)
		}

		cfg = conf
	})
}

func GetConfig() *Config {
	return cfg
}

func Unmarshal(rawVal interface{}) error {
	return viper.Unmarshal(rawVal)
}

func ReadConfig(path string) (*Config, error) {
	conf := &Config{}

	viper.SetDefault("config.path", path)

	err := viper.BindEnv("config.path", "CONFIG_PATH")
	if err != nil {
		log.Printf("warning: %s \n", err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err = viper.MergeInConfig(); err != nil {
		log.Printf("warning: %s \n", err)
		return nil, err
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.Unmarshal(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
