package texts

import "fmt"

const TalkToMeHere = "Поговори мне тут..."

const ThisCommandForTemporaryMute = "Данная команда предназначена для временного mute.\n" +
	"m - минуты, h - часы, d - дни.\nПример использования: /tmute @user 1h 30m"

func UserIsMuted(name string) string {
	return fmt.Sprintf("Пользователь %s теперь будет молчать.", name)
}