package deeprefactoringbot_test

import (
	"github.com/sirupsen/logrus"
	"github.com/deeprefactoring/deeprefactoring-bot"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestApplicationConfig(t *testing.T) {
	validRaw := "log_level: error"
	invalidRaw := "log_level: xxx"

	t.Run("invalid level raises error", func(t *testing.T) {
		var config deeprefactoringbot.ApplicationConfig

		err := yaml.Unmarshal([]byte(invalidRaw), &config)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "can't parse log level xxx")
	})

	t.Run("valid level", func(t *testing.T) {
		var config deeprefactoringbot.ApplicationConfig

		err := yaml.Unmarshal([]byte(validRaw), &config)
		assert.NoError(t, err)
		assert.Equal(t, config.LogLevel, logrus.ErrorLevel)
	})
}
