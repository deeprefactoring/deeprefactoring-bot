package main

import (
	"flag"
	"github.com/Sirupsen/logrus"
	"os"
	"fmt"
	"log"
	"github.com/deeprefactoring/deeprefactoring-bot"
)

var configPath = flag.String("config", "config.yml", "configuration path")
var logLevel = flag.String("log-level", "info", "logging level")

func init() {
	flag.Parse()
}

func initLogger(logLevel string) error {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("can't parse log level %s", logLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(level)
	logrus.SetOutput(os.Stderr)

	return nil
}

func main() {
	err := initLogger(*logLevel)
	if err != nil {
		log.Fatal(err)
	}

	config, err := deeprefactoringbot.NewConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	service, err := deeprefactoringbot.NewService(config.Telegram.ApiKey)
	if err != nil {
		log.Fatal(err)
	}

	service.Listen()
}
