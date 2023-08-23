package config

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Listen struct {
		Type   string `yaml:"type"`
		BindIP string `yaml:"bind_ip"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
}

var once sync.Once

func GetConfigYml() *Config {
	instance := &Config{}
	once.Do(func() {
		if err := cleanenv.ReadConfig("config/config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Info(help)
		}
	})
	return instance
}
