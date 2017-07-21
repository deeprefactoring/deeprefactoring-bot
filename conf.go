package deeprefactoringbot

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"fmt"
)

type TelegramConfig struct {
	ApiKey string `yaml:"api_key"`
}

type Config struct {
	Telegram TelegramConfig `yaml:"telegram"`
}

func NewConfig(path string) (*Config, error) {
	var config Config

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s, err: %s", path, err)
	}
	err = yaml.Unmarshal(contents, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal file %s, err: %s", path, err)
	}

	return &config, nil
}
