package bot

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type TelegramConfig struct {
	ApiKey string `yaml:"api_key"`
}

type ApplicationConfig struct {
	LogLevel logrus.Level
}

func (c *ApplicationConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	aux := &struct {
		LogLevel string `yaml:"log_level"`
	}{}

	err := unmarshal(&aux)
	if err != nil {
		return err
	}

	level, err := logrus.ParseLevel(aux.LogLevel)
	if err != nil {
		return fmt.Errorf("can't parse log level %s", aux.LogLevel)
	}

	c.LogLevel = level
	return nil
}

type Config struct {
	Telegram    TelegramConfig    `yaml:"telegram"`
	Application ApplicationConfig `yaml:"application"`
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
