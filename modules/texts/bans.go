package texts

const YouAreNotAdministrator = "Вы не администратор."

const YouCanNotToDoSomethingWithUsers = "Вы не не имеете права что-то делать с пользователями."

const ICanNotToDoItWithAdministrator = "Я не могу сделать это с администратором."

const UserIsInTheChat = "Этот пользователь в данный момент в чате."

const UserIsNotInTheChat = "Этого пользователя нет в чате."

func UserIsBanned(name string) string {
	return "Пользователь " + name + " забанен!"
}

func UserIsUnbanned(name string) string {
	return "Пользователь " + name + " разбанен!"
}

func UserIsKicked(name string) string {
	return "Пользователь " + name + " кикнут!"
}