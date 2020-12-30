package texts

import "fmt"

const UserIsInTheChat = "Этот пользователь в данный момент в чате."

func UserIsBanned(name string) string {
	return fmt.Sprintf("Пользователь %s забанен!", name)
}

func UserIsUnbanned(name string) string {
	return fmt.Sprintf("Пользователь %s разбанен!", name)
}

func UserIsKicked(name string) string {
	return fmt.Sprintf("Пользователь %s кикнут!", name)
}

func UserIsPromoted(name string) string {
	return fmt.Sprintf("Пользователь %s получил права администратора.", name)
}

func UserIsDemoted(name string) string {
	return fmt.Sprintf("Пользователь %s лишен прав администратора.", name)
}