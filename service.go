package deeprefactoringbot

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"math/rand"
	"strings"
	"time"
)

type Service struct {
	bot    *tgbotapi.BotAPI
	logger *logrus.Entry
}

func NewService(apiKey string) (*Service, error) {
	logger := logrus.WithField("name", "telegram.Service")

	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		logger.WithError(err).Error("failed to connect api")
		return nil, fmt.Errorf("Fails to connect API, err: %s", err)
	}

	logger.Info("Authorized")

	return &Service{logger: logger, bot: bot}, nil
}

func (s *Service) Listen() {
	updateConfig := tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   0,
		Timeout: 60,
	}

	// no error it returns
	updates, _ := s.bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		s.logger.WithField("obj", update).Debug("new update")
		s.handleUpdate(&update)
	}
}

func (s *Service) handleUpdate(update *tgbotapi.Update) {
	message := update.Message
	if message == nil {
		s.logger.Debug("nil message")
		return
	}

	if message.IsCommand() {
		command := message.Command()
		if command == "greeting" {
			s.Greeting(update)
		} else {
			s.logger.WithFields(logrus.Fields{
				"command": message.Command(),
				"message": message,
			}).Warn("unknown command")
		}
	}
	// doing nothing on non commands
}

func (s *Service) Greeting(update *tgbotapi.Update) error {

	messages := make([]string, 0)
	messages = append(messages,
		"{username}, деплоить докером не бросим",
		"Питон лучше руби",
		"Вечер в коворкинг, {username}, деплой в радость, ролбек в сладость",
		"Привет {username}",
	)

	r := rand.New(rand.NewSource(time.Now().Unix()))
	text := messages[r.Intn(len(messages))]

	text = strings.Replace(text, "{username}", fmt.Sprintf("@%s", update.Message.From.UserName), 1)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	s.bot.Send(msg)

	return nil
}
