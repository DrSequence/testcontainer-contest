package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Address    string `yaml:"address"`
		Username   string `yaml:"user"`
		Password   string `yaml:"pass"`
		Database   string `yaml:"database"`
		Collection string `yaml:"collection"`
	} `yaml:"database"`
	Cache struct {
		Address string        `yaml:"address"`
		Exp     time.Duration `yaml:"exp"`
		Pass    string        `yaml:"pass"`
	} `yaml:"cache"`
}

func ReadConfig() Config {
	f, err := os.Open("config/app.yaml")
	if err != nil {
		log.Fatalf("error while oping file. %s", err)
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("error while parsing config. %s", err)
	}

	return cfg
}
