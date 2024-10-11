package config

import (
	"gopkg.in/yaml.v3"
	"log"
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

func ReadConfig(cfgPath string) Config {
	file, err := os.Open(cfgPath)
	if err != nil {
		log.Fatal(err) // return err
	}

	defer func() {
		err := file.Close() // you can just defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	cfg := Config{}
	err = yaml.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Fatal(err) // return err
	}

	return cfg
}
