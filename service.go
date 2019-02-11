package deeprefactoringbot

import (
	"fmt"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/sirupsen/logrus"
)

// Generic (joke) bot API interface to use in tests,
// consider tgbotapi as implementation
type BotAPI interface {
	GetUpdatesChan(config tgbotapi.UpdateConfig) (tgbotapi.UpdatesChannel, error)
	Send(c tgbotapi.Chattable) (tgbotapi.Message, error)
}

type Service struct {
	bot    BotAPI
	logger *logrus.Entry
	msg    MessageProvider
}

func NewServiceFromTgbotapi(apiKey string, msg MessageProvider) (*Service, error) {
	logger := logrus.WithField("name", "telegram.Service")

	bot, err := tgbotapi.NewBotAPIWithClient(
		apiKey,
		&http.Client{
			Timeout: time.Duration(60) * time.Second,
		},
	)
	if err != nil {
		logger.WithError(err).Error("failed to connect api")
		return nil, fmt.Errorf("fails to connect API, err: %s", err)
	}

	logger.Info("Authorized")

	return NewService(bot, msg, logger), nil
}

func NewService(bot BotAPI, msg MessageProvider, logger *logrus.Entry) *Service {
	return &Service{logger: logger, bot: bot, msg: msg}
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
		s.HandleUpdate(&update)
	}
}

func (s *Service) HandleUpdate(update *tgbotapi.Update) {
	s.logger.WithFields(logrus.Fields{
		"Update":            update,
		"ChannelPost":       update.ChannelPost,
		"Message":           update.Message,
		"EditedChannelPost": update.EditedChannelPost,
	}).Debug("new update")

	message := update.Message
	if message == nil {
		s.logger.Debug("nil message")
		return
	}

	if message.IsCommand() {
		command := message.Command()
		switch command {
		case "greeting":
			s.Greeting(update, update.Message.From.String())
		case "nextmeetup":
			s.NextMeetup(update)
		case "roll":
			s.RollMessage(update)
		case "stop":
			s.Hammertime(update)
		default:
			s.logger.WithFields(logrus.Fields{
				"command": message.Command(),
				"message": message,
			}).Warn("unknown command")
		}
	}

	if update.Message.NewChatMembers != nil {
		for _, user := range *update.Message.NewChatMembers {
			s.Greeting(update, user.String())
		}
	}

	if update.Message.LeftChatMember != nil {
		s.GoAwayMessage(update, update.Message.LeftChatMember.String())
	}
}

func (s *Service) Send(update *tgbotapi.Update, text string) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	s.bot.Send(msg)
	return nil
}

func (s *Service) Greeting(update *tgbotapi.Update, username string) error {
	text := ReplaceUsername(s.msg.GetGreeting(), username)
	return s.Send(update, text)
}

func (s *Service) GoAwayMessage(update *tgbotapi.Update, username string) error {
	text := ReplaceUsername(s.msg.GetCurse(), username)
	return s.Send(update, text)
}

func (s *Service) NextMeetup(update *tgbotapi.Update) error {
	text := NextMeetupInfo()
	return s.Send(update, text)
}

func (s *Service) RollMessage(update *tgbotapi.Update) error {
	text := s.msg.GetRoll()
	return s.Send(update, text)
}

func (s *Service) Hammertime(update *tgbotapi.Update) error {
	text := HammertimeInfo()
	return s.Send(update, text)
}

func (s *Service) GetMessage() MessageProvider {
	return s.msg
}
