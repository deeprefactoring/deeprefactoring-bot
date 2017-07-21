# Introduction
This is a community bot of Глубокий Рефакторинг (Deep Refactoring). It can handle various things including producing sarcastic comments on every [@achikin][1] post (not enabled by default).

# Development guideline
## Create bot api_key
* In telegram open dialog with [@BotFather][2]
* Follow instructions there to create bot and obtain api_key

## Code preparation
* Install [Go 1.7][3]
* Download code to `$GOPATH/src/github.com/deeprefactoring/deeprefactoring-bot` folder:
```bash
go get github.com/deeprefactoring/deeprefactoring-bot
```
* Copy `config.yml.example` to `config.yml`
* Use api_key from previous steps in `config.yaml`
* run by calling `go run ./cmd/app/main.go`

[1]: https://github.com/achikin
[2]: https://telegram.me/BotFather
[3]: https://golang.org/dl
