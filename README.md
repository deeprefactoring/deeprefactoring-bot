[![CircleCI](https://circleci.com/gh/deeprefactoring/deeprefactoring-bot/tree/master.svg?style=svg&circle-token=0e2f1cd5497fa9397ce7905df9fe92a2ad4ca86a)](https://circleci.com/gh/deeprefactoring/deeprefactoring-bot/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/deeprefactoring/deeprefactoring-bot)](https://goreportcard.com/report/github.com/deeprefactoring/deeprefactoring-bot)

# Introduction
This is a community bot of Глубокий Рефакторинг (Deep Refactoring). It can handle various things including producing sarcastic comments on every @achikin post (not enabled by default).

# Development guideline
## create bot api_key
* In telegram open dialog with @botfather
* Follow instructions there to create bot and obtain api_key
## Code preparation
* Install Go 1.7
* Clone repository to GOPATH/src/github.com/deeprefactoring/deeprefactoring-bot folder
* Copy config.yml.example to config.yml
* Use api_key from previous steps in config.yaml
* run by calling go run ./cmd/app/main.go
