package texts

const TalkToMeHere = "Поговори мне тут..."

const ThisCommandForTemporaryMute = "Данная команда предназначена для временного mute."

func UserIsMuted(name string) string {
	return "Пользователь " + name + " теперь будет молчать."
}