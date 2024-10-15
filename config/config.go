package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	Config struct {
		DBConfig DBConfig `yaml:"db_config"`
		Server   Server   `yaml:"server"`
	}
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	}
	DBConfig struct {
		Migrates string `yaml:"migrates"`
		Driver   string `yaml:"driver"`
		Postgres Connector
	}
	Connector struct {
		ConnectionDSN string `yaml:"connection_dsn"`
	}
)

func (c Connector) DSN() (string, error) {
	return c.ConnectionDSN, nil
}

func ReadConfig(cfgPath string) (Config, error) {
	file, err := os.Open(cfgPath)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	cfg := Config{}
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, err
}
