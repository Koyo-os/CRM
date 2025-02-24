package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct{
	Port uint64 `yaml:"port"`
	Host string `yaml:"host"`
	MongoURL string `yaml:"mongo_url"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil{
		return nil, err
	}

	body, err := io.ReadAll(file)
	if err != nil{
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(body, &cfg)
	return &cfg, err
}