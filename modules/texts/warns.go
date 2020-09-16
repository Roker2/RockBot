package texts

import "strconv"

const ICanNotGiveWarnToAdministrator = "Я не могу дать предупреждение администратору."

func WarnsQuantityOfUser(FirstName string, quantity int, maxQuantity int) string {
	return "Количество предупреждений у " +FirstName + ": " + strconv.Itoa(quantity) + "/" + strconv.Itoa(maxQuantity)
}

func NewWarnsQuantity(quantity string) string {
	return "Новое количество максимальных предупреждений: " + quantity + "."
}

const WritePleaseInteger = "Введите пожалуйста число."

const AdministratorAlwaysIsClean = "У администраторов всегда карма чистая. По крайней мере здесь."

func UserDoesNotHaveWarns(FirstName string) string {
	return "У пользователя " + FirstName + " очищена карма."
}

const RemoveWarn = "Убрать предупреждение"

const WarnWasRemoved = "Предупреждение убрано."