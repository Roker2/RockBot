package texts

const DescriptionOfRulesCommand = "Эта комманда позволяет установить правила."

const NewRulesWereAdded = "Правила установлены! Вы можете посмотреть их с помощью команды /rules."

func ErrorOfGettingRules(errText string) string {
	return "Ошибка получения правил.\n" + errText
}