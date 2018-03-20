package deeprefactoringbot

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var GreetingMessages = []string{
	"{username}, деплоить докером не бросим",
	"Питон лучше руби",
	"Руби лучше питон",
	"Элексир лучше питона и руби",
	"Вечер в коворкинг, {username}, деплой в радость, ролбек в сладость",
	"Привет, {username}",
	"{username}, опиши свой проект версией дебиана",
	"{username}, сколько ты принес прибыли компании сегодня?",
	"raise Exception('hello {username}')",
	"Лучший воронежский чат о рефакторинге привествует тебя, {username}",
	"1 1111 11111111 11111111",
	"Какой JS фреймворк ты предпочитаешь, {username}?",
}

var CurseMessages = []string{
	"{username}, мы не будем скучать",
	"наверное {username} фронтэндер",
	"Press F to Pay Respects for {username}",
	"Мы тебя тоже не любим",
}

var RollMessages = []string{
	"Живые форки php4",
	"Переходные процессы в проектах однодневках",
	"Статическая типизация аналогов баша",
	"Перспективы перехода на subservion в рамках энтерпрайз проекта",
	"Логин от рута, преимущества и руткиты",
	"Санный или ссаный .net",
	"Шифрование методом редупликации, секретики-хуетики",
	"Самый быстрый способ уйти в swap",
	"Сколько в вашей команде женщин за 40 без кошек?",
	"Жать ли руку руби программистам?",
	"Преимущество использования widows для хостинга статичных файлов",
	"Почему питон делится на 6?",
	"Проектный менеджер, что ты несешь?",
	"Лучший шаблон шаблона на с++ 2017",
	"Ранний уход на пенсию по причине создания js фреймворка",
	"Сколько раз ты ронял базу в этом месяце (опрос)",
	"Монитора два, а продакшен сервер один, за и против",
	"Как лучше посылать логи в /dev/null",
	"Использование mutex в языках с GIL",
	"Веб 1.0 на перл, перекличка работников Яндекса",
	"Когда будет новый пост от Чистякова в его канале?",
	"Дженерики в баше, хватит это терпеть",
	"Документация проекта очевидными коментариями",
	"Безключевой доступ в AWS",
	"Jira и другие способы почувстовать себя плохо",
	"Выход из vim народными средствами",
	"Рациональны предложения менеджеру, которые ты никогда не озвучишь",
	"Расчет 6%% на Spring Boot или патент?",
	"Че, пацаны, фронтенд?"
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
