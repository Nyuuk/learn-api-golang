package helpers

import (
	"gopkg.in/yaml.v3"
	"os"
)

type ConfigHelper struct {
	Database struct {
		Postgres struct {
			Host     string `yaml:"host"`
			Port     string `yaml:"port"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Database string `yaml:"database"`
		} "yaml:postgres"
	} "yaml:database"

	App struct {
		Port string `yaml:"port"`
		HelloMsg string `yaml:"hello_msg"`
		Name string `yaml:"name"`
	} "yaml:app"
}

func LoadConfig(filename string) (*ConfigHelper, error) {
	config := &ConfigHelper{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}