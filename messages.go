package deeprefactoringbot

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var GreetingMessages = []string{
	"Руби лучше питон",
	"Элексир лучше питона и руби",
	"Какой JS фреймворк ты предпочитаешь, {username}?",
	"А твой PM знает об этом?",
	"Возможно, мы знаем кое-что о тебе?",
	"Люди всё время меня спрашивают: знаю ли я {username}?",
}

var CurseMessages = []string{
	"{username} ушел в реальный мир",
}

var RollMessages = []string{
	"Расчет 6%% на Spring Boot или патент?",
	"Че, пацаны, фронтенд?",
	"JavaScript — всему голова",
	"Жизнь похожа на коробку NPM пакетов",
}

const nextMeetupMessage = "Анонс следующего митапа: http://deeprefactoring.ru/meetup/next/"
const hammertimeMessage = "https://ci.memecdn.com/2501287.gif"

var randomiser = rand.New(rand.NewSource(time.Now().Unix()))

func randomMessageFromSlice(slice []string) string {
	return slice[randomiser.Intn(len(slice))]
}

func ReplaceUsername(text, username string) string {
	return strings.Replace(text, "{username}", fmt.Sprintf("@%s", username), -1)
}

func RandomGreeting(username string) string {
	text := randomMessageFromSlice(GreetingMessages)
	return ReplaceUsername(text, username)
}

func RandomCurse(username string) string {
	text := randomMessageFromSlice(CurseMessages)
	return ReplaceUsername(text, username)
}

func NextMeetupInfo() string {
	return nextMeetupMessage
}

func HammertimeInfo() string {
	return hammertimeMessage
}

func RollMessage() string {
	return randomMessageFromSlice(RollMessages)
}
