package texts

import "strconv"

const ICanNotGiveWarnToAdministrator = "Я не могу дать предупреждение администратору."

const WritePleaseInteger = "Введите пожалуйста число."

const AdministratorAlwaysIsClean = "У администраторов всегда карма чистая. По крайней мере здесь."

const RemoveWarn = "Убрать предупреждение"

const WarnWasRemoved = "Предупреждение убрано."

func WarnsQuantityOfUser(FirstName string, quantity int, maxQuantity int) string {
	return "Количество предупреждений у " + FirstName + ": " + strconv.Itoa(quantity) + "/" + strconv.Itoa(maxQuantity)
}

func NewWarnsQuantity(quantity string) string {
	return "Новое количество максимальных предупреждений: " + quantity + "."
}

func UserDoesNotHaveWarns(FirstName string) string {
	return "У пользователя " + FirstName + " очищена карма."
}