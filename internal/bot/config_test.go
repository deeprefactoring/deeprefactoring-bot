package bot_test

import (
	"testing"

	deepbot "github.com/deeprefactoring/deeprefactoring-bot/internal/bot"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestApplicationConfig(t *testing.T) {
	validRaw := "log_level: error"
	invalidRaw := "log_level: xxx"

	t.Run("invalid level raises error", func(t *testing.T) {
		var config deepbot.ApplicationConfig

		err := yaml.Unmarshal([]byte(invalidRaw), &config)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "can't parse log level xxx")
	})

	t.Run("valid level", func(t *testing.T) {
		var config deepbot.ApplicationConfig

		err := yaml.Unmarshal([]byte(validRaw), &config)
		assert.NoError(t, err)
		assert.Equal(t, config.LogLevel, logrus.ErrorLevel)
	})
}
