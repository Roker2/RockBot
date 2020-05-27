package texts

const TalkToMeHere = "Поговори мне тут..."

const ThisCommandForTemporaryMute = "Данная команда предназначена для временного mute.\nm - минуты, h - часы, d - дни.\nПример использования: /mute @user 1h 30m"

func UserIsMuted(name string) string {
	return "Пользователь " + name + " теперь будет молчать."
}