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

var randomiser = rand.New(rand.NewSource(time.Now().Unix()))

func ReplaceUsername(text, username string) string {
	return strings.Replace(text, "{username}", fmt.Sprintf("@%s", username), -1)
}

func RandomGreeting(username string) string {
	text := greetingMessages[randomiser.Intn(len(greetingMessages))]
	return ReplaceUsername(text, username)
}
