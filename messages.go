package deeprefactoringbot

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var greetingMessages = []string{
	"{username}, деплоить докером не бросим",
	"Питон лучше руби",
	"Вечер в коворкинг, {username}, деплой в радость, ролбек в сладость",
	"Привет {username}",
}

var curseMessages = []string{
	"{username}, мы не будем скучать",
	"наверное {username} фронтэндер",
	"Press F to Pay Respects for {username}",
	"Мы тебя тоже не любим",
}

var randomiser = rand.New(rand.NewSource(time.Now().Unix()))

func ReplaceUsername(text, username string) string {
	return strings.Replace(text, "{username}", fmt.Sprintf("@%s", username), -1)
}

func RandomGreeting(username string) string {
	text := greetingMessages[randomiser.Intn(len(greetingMessages))]
	return ReplaceUsername(text, username)
}

func RandomCurse(username string) string {
	text := curseMessages[randomiser.Intn(len(curseMessages))]
	return ReplaceUsername(text, username)
}
