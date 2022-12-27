package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database DbConfig `yaml:"database_config" binding:"required"`
}

type DbConfig struct {
	Username string     `yaml:"username" binding:"required"`
	Password string     `yaml:"password" binding:"required"`
	Host     string     `yaml:"hostname" binding:"required"`
	Port     string     `yaml:"port" binding:"required"`
	DbName   string     `yaml:"database" binding:"required"`
	Tables   TableNames `yaml:"table_names" binding:"required"`
}

type TableNames struct {
	Shorteners string `yaml:"shorteners" binding:"required"`
}

func LoadConfig(path string) (Config, error) {
	var (
		cfg Config
		err error
	)

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}
