package config

import (
	"github.com/spf13/viper"
	"time"
)

type ServerConfig struct {
	Port string `yaml:"port"`
}

type DbConfig struct {
	DataSourceName string `yaml:"dataSourceName"`
	Migrate        struct {
		Enable bool `yaml:"enable"`
	} `yaml:"migrate"`
	Pool struct {
		MaxOpen     int           `yaml:"maxOpen"`
		MaxIdle     int           `yaml:"maxIdle"`
		MaxLifeTime time.Duration `yaml:"maxLifeTime"`
	}
}

type Config struct {
	Server ServerConfig `yaml:"server"`
	Db     DbConfig     `yaml:"db"`
}

func LoadConfig() (Config, error) {
	vp := viper.New()
	var config Config
	vp.AddConfigPath("././configs/")
	vp.SetConfigName("config")
	vp.SetConfigType("yaml")

	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}

	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
