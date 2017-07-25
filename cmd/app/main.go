package main

import (
	"flag"
	"fmt"
	"github.com/deeprefactoring/deeprefactoring-bot"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var buildVersion, buildDate string

var Arguments struct {
	ConfigPath string
	Version    bool
}

func init() {
	flag.StringVar(&Arguments.ConfigPath, "config", "config.yml", "configuration path")
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

	initLogger(config.Application.LogLevel)

	service, err := deeprefactoringbot.NewService(config.Telegram.ApiKey)
	if err != nil {
		log.Fatal(err)
	}

	service.Listen()
}
