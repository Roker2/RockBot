package texts

const UserIsInTheChat = "Этот пользователь в данный момент в чате."

func UserIsBanned(name string) string {
	return "Пользователь " + name + " забанен!"
}

func UserIsUnbanned(name string) string {
	return "Пользователь " + name + " разбанен!"
}

func UserIsKicked(name string) string {
	return "Пользователь " + name + " кикнут!"
}

func UserIsPromoted(name string) string {
	return "Пользователь " + name + " получил права администратора."
}

func UserIsDemoted(name string) string {
	return "Пользователь " + name + " лишен прав администратора."
}