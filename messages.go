package deeprefactoringbot

import (
	"fmt"
	"strings"
)

const nextMeetupMessage = "Анонс следующего митапа: http://deeprefactoring.ru/meetup/next/"
const hammertimeMessage = "https://ci.memecdn.com/2501287.gif"

// MessageProvider common chat message repo interface
type MessageProvider interface {
	// GetGreeting returns a greeting message
	GetGreeting() string
	// GetCurse returns a farewell message
	GetCurse() string
	// GetRoll returns a topic message
	GetRoll() string
}

func ReplaceUsername(text, username string) string {
	return strings.Replace(text, "{username}", fmt.Sprintf("@%s", username), -1)
}

func NextMeetupInfo() string {
	return nextMeetupMessage
}

func HammertimeInfo() string {
	return hammertimeMessage
}
