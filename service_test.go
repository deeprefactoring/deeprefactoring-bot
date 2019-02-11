package deeprefactoringbot_test

import (
	"testing"

	deeprefactoringbot "github.com/deeprefactoring/deeprefactoring-bot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

type FakeBot struct {
	SentMessages []tgbotapi.Chattable
	Channel      chan tgbotapi.Update
}

func (bot *FakeBot) GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error) {
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

type FakeMessage struct{}

func (f *FakeMessage) GetGreeting() string { return "Greeting" }
func (f *FakeMessage) GetCurse() string    { return "Curse" }
func (f *FakeMessage) GetRoll() string     { return "Roll" }

func ServiceWithLogger(bot deeprefactoringbot.BotAPI) (*deeprefactoringbot.Service, *test.Hook) {
	logger, hook := test.NewNullLogger()
	logger.Level = logrus.DebugLevel

	service := deeprefactoringbot.NewService(bot, &FakeMessage{}, logger.WithField("name", "logger"))

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
			Entities: &[]tgbotapi.MessageEntity{
				{Type: "bot_command", Offset: 0, Length: 11},
			},
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
			Entities: &[]tgbotapi.MessageEntity{
				{Type: "bot_command", Offset: 0, Length: 5},
			},
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
			Entities: &[]tgbotapi.MessageEntity{
				{Type: "bot_command", Offset: 0, Length: 5},
			},
			Chat: &tgbotapi.Chat{
				ID: 1,
			},
			Text: "/roll",
			From: &user,
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Equal(
		t,
		service.GetMessage().GetRoll(),
		bot.LastMessageConfig().Text,
	)
}

func TestService_Greeting(t *testing.T) {
	bot := &FakeBot{}

	service, _ := ServiceWithLogger(bot)

	user := tgbotapi.User{UserName: "Hoi"}

	service.HandleUpdate(&tgbotapi.Update{
		Message: &tgbotapi.Message{
			Entities: &[]tgbotapi.MessageEntity{
				{Type: "bot_command", Offset: 0, Length: 9},
			},
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
		service.GetMessage().GetGreeting(),
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
			NewChatMembers: &[]tgbotapi.User{user},
		},
	})

	assert.Equal(t, len(bot.SentMessages), 1)
	assert.Contains(
		t,
		service.GetMessage().GetGreeting(),
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
		service.GetMessage().GetCurse(),
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
