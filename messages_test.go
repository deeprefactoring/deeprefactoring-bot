package deeprefactoringbot_test

import (
	"github.com/deeprefactoring/deeprefactoring-bot"
	"github.com/stretchr/testify/assert"
	"testing"
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
		actual := deeprefactoringbot.ReplaceUsername(c.template, c.username)
		assert.Equal(t, c.expected, actual)
	}
}

func TestHammertime(t *testing.T) {
	assert.Equal(t, deeprefactoringbot.HammertimeInfo(), "https://ci.memecdn.com/2501287.gif")
}

func TestRandomCurse(t *testing.T) {
	assert.NotEqual(t, deeprefactoringbot.RandomCurse("x"), "")
}

func TestRandomGreeting(t *testing.T) {
	assert.NotEqual(t, deeprefactoringbot.RandomGreeting("x"), "")
}
