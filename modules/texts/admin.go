package texts

const PleaseReplyToTheMessageYouWantToPin = "Ответьте пожалуйста на сообщение, которое Вы хотите закрепить."

const ThisChatIsPrivateICanNotToPinMessage = "Данный чат приватный, в приватных чатах я не могу закрепить сообщение."

const ThisChatIsPrivateICanNotToUnpinMessage = "Данный чат приватный, в приватных чатах я не могу открепить сообщение."

const PleaseReplyToTheMessageOfThePersonYouWantToGrantAdministratorRightsToOrEnterTheirID = "Ответьте пожалуйста на сообщение того, кому Вы хотите выдать права администратора, или введите его ID."

const PleaseReplyToTheMessageOfThePersonYouWantToRemoveAdministratorRightsToOrEnterTheirID = "Ответьте пожалуйста на сообщение того, кому Вы хотите убрать права администратора, или введите его ID."

const PurgeCompleted = "Очистка завершена. Сообщение удалится через 5 секунд."

const AllUserCommandsAreDisabled = "Все пользовательские команды отключены."

const AllUserCommandsAreEnabled = "Все пользовательские команды включены."

const YouDidNotWriteAnyUserCommands = "Вы не написали ни одной пользовательской команды."

func DisabledUserCommandsList(disabledCommands string) string {
	return "Отключены следующие пользовательские команды: " + disabledCommands
}