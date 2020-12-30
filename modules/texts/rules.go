package texts

import "fmt"

const DescriptionOfRulesCommand = "Эта комманда позволяет установить правила."

const NewRulesWereAdded = "Правила установлены! Вы можете посмотреть их с помощью команды /rules."

func ErrorOfGettingRules(errText string) string {
	return fmt.Sprintf("Ошибка получения правил.\n%s", errText)
}