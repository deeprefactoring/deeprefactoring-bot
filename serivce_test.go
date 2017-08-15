package deeprefactoringbot_test

import (
	"github.com/deeprefactoring/deeprefactoring-bot"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FakeBot struct {
	SentMessages []tgbotapi.Chattable
	Channel      chan tgbotapi.Update
}

func (bot *FakeBot) GetUpdatesChan(config tgbotapi.UpdateConfig) (<-chan tgbotapi.Update, error) {
	return bot.Channel, nil
}

func (bot *FakeBot) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	bot.SentMessages = append(bot.SentMessages, c)
	return tgbotapi.Message{}, nil
}

func (bot *FakeBot) LastChattable() tgbotapi.Chattable {
	return bot.SentMessages[len(bot.SentMessages)-1]
}

func (bot *FakeBot) LastMessageConfig() tgbotapi.MessageConfig {
	return bot.LastChattable().(tgbotapi.MessageConfig)
}

func ServiceWithLogger(bot deeprefactoringbot.BotAPI) (*deeprefactoringbot.Service, *test.Hook) {
	logger, hook := test.NewNullLogger()
	logger.Level = logrus.DebugLevel

	service := deeprefactoringbot.NewService(bot, logger.WithField("name", "logger"))

	return service, hook
}

func TestService_Listen(t *testing.T) {
	ch := make(chan tgbotapi.Update, 100)

	bot := &FakeBot{Channel: ch}

	user := tgbotapi.User{UserName: "user1"}

	ch <- tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "foo",
			From: &user,
		},
	}

	ch <- tgbotapi.Update{
		Message: &tgbotapi.Message{
			Text: "bar",
			From: &user,
		},
	}

	service, hook := ServiceWithLogger(bot)
	go close(ch)
	service.Listen()

	entries := hook.AllEntries()
	assert.Equal(t, 2, len(entries))

	fooMessage := entries[0]
	assert.Equal(t, "new update", fooMessage.Message)
	assert.Equal(t, "foo", fooMessage.Data["Message"].(*tgbotapi.Message).Text)

	barMessage := entries[1]
	assert.Equal(t, "new update", barMessage.Message)
	assert.Equal(t, "bar", barMessage.Data["Message"].(*tgbotapi.Message).Text)
}

func TestService_HandleUpdateFiltered(t *testing.T) {

	cases := []struct {
		name     string
		updates  []tgbotapi.Update
		expected []string
	}{
		{"empty messages", []tgbotapi.Update{}, []string{}},
		{"stripped update", []tgbotapi.Update{{}}, []string{}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			bot := &FakeBot{}

			service, _ := ServiceWithLogger(bot)
			for _, update := range c.updates {
				service.HandleUpdate(&update)
			}

			assert.Equal(t, len(bot.SentMessages), len(c.expected))
		})
	}
}

func TestService_NextMeetup(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)
	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			Text: "/nextmeetup",
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(t, bot.LastMessageConfig().Text, "Анонс следующего митапа")
}
func TestService_Hammertime(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)
	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			Text: "/stop",
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(t, bot.LastMessageConfig().Text, "https://ci.memecdn.com/2501287.gif")
}

func applyUsername(slice []string, username string) []string {
	res := make([]string, len(slice))

	for _, value := range slice {
		res = append(res, deeprefactoringbot.ReplaceUsername(value, username))
	}

	return res
}

func TestService_RollMessage(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)

	user := tgbotapi.User{UserName: "Hoi"}

	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			Text: "/roll",
			From: &user,
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(
		t,
		applyUsername(deeprefactoringbot.RollMessages, user.UserName),
		bot.LastMessageConfig().Text,
	)
}

func TestService_Greeting(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)

	user := tgbotapi.User{UserName: "Hoi"}

	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			Text: "/greeting",
			From: &user,
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(
		t,
		applyUsername(deeprefactoringbot.GreetingMessages, user.UserName),
		bot.LastMessageConfig().Text,
	)
}

func TestService_Greeting2(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)

	user := tgbotapi.User{UserName: "Hoi"}

	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			NewChatMember: &user,
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(
		t,
		applyUsername(deeprefactoringbot.GreetingMessages, user.UserName),
		bot.LastMessageConfig().Text,
	)
}

func TestService_GoAwayMessage(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)

	user := tgbotapi.User{UserName: "Hoi"}

	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			LeftChatMember: &user,
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(
		t,
		applyUsername(deeprefactoringbot.CurseMessages, user.UserName),
		bot.LastMessageConfig().Text,
	)
}

func TestService_Undefined(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)

	user := tgbotapi.User{UserName: "Hoi"}

	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			Text: "/undefined",
			From: &user,
		},
	})

	assert.Equal(t, len(bot.SentMessages), 0)
}
