package texts

import (
	"fmt"
)

const ICanNotGiveWarnToAdministrator = "Я не могу дать предупреждение администратору."

const WritePleaseInteger = "Введите пожалуйста число."

const AdministratorAlwaysIsClean = "У администраторов всегда карма чистая. По крайней мере здесь."

const RemoveWarn = "Убрать предупреждение"

const WarnWasRemoved = "Предупреждение убрано."

func WarnsQuantityOfUser(FirstName string, quantity int, maxQuantity int) string {
	return fmt.Sprintf("Количество предупреждений у %s: %d/%d", FirstName, quantity, maxQuantity)
}

func NewWarnsQuantity(quantity string) string {
	return fmt.Sprintf("Новое количество максимальных предупреждений: %s.", quantity)
}

func UserDoesNotHaveWarns(FirstName string) string {
	return fmt.Sprintf("У пользователя %s очищена карма.", FirstName)
}