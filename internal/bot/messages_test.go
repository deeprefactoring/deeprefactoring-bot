package bot_test

import (
	"testing"

	"github.com/deeprefactoring/deeprefactoring-bot/internal/bot"

	"github.com/stretchr/testify/assert"
)

func TestReplaceUsername(t *testing.T) {

	cases := []struct {
		template string
		username string
		expected string
	}{
		{"hello {username}!", "name", "hello @name!"},
		{"hello {user}!", "name", "hello {user}!"},
	}

	for _, c := range cases {
		actual := bot.ReplaceUsername(c.template, c.username)
		assert.Equal(t, c.expected, actual)
	}
}

func TestHammertime(t *testing.T) {
	assert.Equal(t, bot.HammertimeInfo(), "https://ci.memecdn.com/2501287.gif")
}
