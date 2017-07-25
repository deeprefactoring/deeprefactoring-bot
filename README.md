[![CircleCI](https://circleci.com/gh/deeprefactoring/deeprefactoring-bot/tree/master.svg?style=svg&circle-token=0e2f1cd5497fa9397ce7905df9fe92a2ad4ca86a)](https://circleci.com/gh/deeprefactoring/deeprefactoring-bot/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/deeprefactoring/deeprefactoring-bot)](https://goreportcard.com/report/github.com/deeprefactoring/deeprefactoring-bot)

# Introduction
This is a community bot of Глубокий Рефакторинг (Deep Refactoring). It can handle various things including producing sarcastic comments on every [@achikin]<sup name="a1">[1](#f1)</sup> post (not enabled by default).

# Development guideline
## Create bot api_key
* In telegram open dialog with [@BotFather]<sup name="a2">[2](#f2)</sup>
* Follow instructions there to create bot and obtain api_key

## Contribution
* Install [Go 1.8+]<sup name="a3">[3](#f3)</sup>
* Download code to `$GOPATH/src/github.com/deeprefactoring/deeprefactoring-bot` folder:
```bash
mkdir -p github.com/deeprefactoring
cd github.com/deeprefactoring
git clone github.com/deeprefactoring/deeprefactoring-bot
```
* Copy `config.yml.example` to `config.yml`
* Use api_key from previous steps in `config.yaml`
* Install dependencies `make deps`
* Build `make build`, produces `deeprefactoring-bot` binary
* Run tests `make test`

<span name="f1">[1]:</span> Famous deep refactoring collaborator Anton Chikin in Telegram [↩](#a1)
<br>
<span name="f2">[2]:</span> https://telegram.me/BotFather [↩](#a2)
<br>
<span name="f3">[3]:</span> https://golang.org/dl [↩](#a3)
