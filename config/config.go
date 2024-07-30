package config

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strings"
)

func loadConfigFile(configFilePath string) error {
	viper.SetConfigFile(configFilePath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))

	if err := viper.ReadInConfig(); err != nil {
		log.Errorf("Can't read config file: %v\n", err)
		return err
	} else {
		log.Infof("Config file loaded")
		return nil
	}
}

func ParseArgs() string {
	var configFilePath string
	flag.StringVar(&configFilePath, "c", "./config.yml", "path to config file location")
	flag.Parse()
	log.Infof("Config file path: %v\n", configFilePath)

	return configFilePath
}

func LoadConfig[T any](configFilePath string) (*T, error) {
	err := loadConfigFile(configFilePath)
	if err != nil {
		return nil, err
	}

	t := new(T)

	if err = viper.Unmarshal(t); err != nil {
		log.Errorf("Can't unmarshal config file: %v\n", err)
		return nil, err
	} else {
		log.Infof("Config file unmarshalled")
	}

	return t, nil
}
