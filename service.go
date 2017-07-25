package deeprefactoringbot

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
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
		s.logger.WithFields(logrus.Fields{
			"update":            update,
			"ChannelPost":       update.ChannelPost,
			"Message":           update.Message,
			"EditedChannelPost": update.EditedChannelPost,
		}).Debug("new update")
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
		switch command {
		case "greeting":
			s.Greeting(update, update.Message.From.UserName)
		case "nextmeetup":
			s.NextMeetup(update)
		case "roll":
			s.RollMessage(update)
		default:
			s.logger.WithFields(logrus.Fields{
				"command": message.Command(),
				"message": message,
			}).Warn("unknown command")
		}
	}

	if update.Message.NewChatMember != nil {
		s.Greeting(update, update.Message.NewChatMember.UserName)
	}

	if update.Message.LeftChatMember != nil {
		s.GoAwayMessage(update, update.Message.LeftChatMember.UserName)
	}
}

func (s *Service) Send(update *tgbotapi.Update, text string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	s.bot.Send(msg)
	return nil
}

func (s *Service) Greeting(update *tgbotapi.Update, username string) error {
	text := RandomGreeting(username)
	return s.Send(update, text)
}

func (s *Service) GoAwayMessage(update *tgbotapi.Update, username string) error {
	text := RandomCurse(username)
	return s.Send(update, text)
}

func (s *Service) NextMeetup(update *tgbotapi.Update) error {
	text := NextMeetupInfo()
	return s.Send(update, text)
}

func (s *Service) RollMessage(update *tgbotapi.Update) error {
	text := RollMessage()
	return s.Send(update, text)
}
