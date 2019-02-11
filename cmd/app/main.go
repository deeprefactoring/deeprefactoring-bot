package main

import (
	"flag"
	"fmt"
	"github.com/deeprefactoring/deeprefactoring-bot"
	"github.com/deeprefactoring/deeprefactoring-bot/internal/message"
	"github.com/sirupsen/logrus"
	"log"
	"os"

	deeprefactoringbot "github.com/deeprefactoring/deeprefactoring-bot"
	"github.com/deeprefactoring/deeprefactoring-bot/internal/message"
	"github.com/sirupsen/logrus"
)

var buildVersion, buildDate string

var Arguments struct {
	ConfigPath  string
	MessagePath string
	Version     bool
}

func init() {
	flag.StringVar(&Arguments.ConfigPath, "config", "config.yml", "configuration path")
	flag.StringVar(&Arguments.MessagePath, "messages", "messages.yml", "messages config path")
	flag.BoolVar(&Arguments.Version, "version", false, "output version information")
	flag.Parse()
}

func initLogger(level logrus.Level) {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stderr)
}

func main() {
	if Arguments.Version {
		fmt.Println("build version:", buildVersion)
		fmt.Println("build date:", buildDate)
		os.Exit(0)
	}

	config, err := deeprefactoringbot.NewConfig(Arguments.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	msg, err := message.NewFileMessage(Arguments.MessagePath)
	if err != nil {
		log.Fatal(err)
	}

	initLogger(config.Application.LogLevel)

	service, err := deeprefactoringbot.NewServiceFromTgbotapi(config.Telegram.ApiKey, msg)
	if err != nil {
		log.Fatal(err)
	}

	service.Listen()
}
