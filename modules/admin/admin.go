package admin

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
	"strconv"
	"strings"
	"time"
)

const userCommands = "randomal welcome ping randomsmmq"

func Pin(bot ext.Bot, u *gotgbot.Update, args []string) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		_, err = bot.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	if !utils.MemberCanPin(bot, u) {
		return nil
	}
	if !utils.BotCanPin(bot, u) {
		return nil
	}
	if u.Message.ReplyToMessage == nil {
		_, err := bot.SendMessage(u.Message.Chat.Id, texts.PleaseReplyToTheMessageYouWantToPin)
		return err
	}
	if u.Message.Chat.Type == "private" {
		_, err := bot.SendMessage(u.Message.Chat.Id, texts.ThisChatIsPrivateICanNotToPinMessage)
		return err
	}
	Notify := true
	if len(args) > 0 {
		Notify = strings.ToLower(args[0]) != "loud"
	}
	Message := bot.NewSendablePinChatMessage(u.Message.Chat.Id, u.Message.ReplyToMessage.MessageId)
	Message.DisableNotification = Notify
	_, err = Message.Send()
	return err
}

func Unpin(bot ext.Bot, u *gotgbot.Update) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		_, err = bot.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	if !utils.MemberCanPin(bot, u) {
		return nil
	}
	if !utils.BotCanPin(bot, u) {
		return nil
	}
	if u.Message.Chat.Type == "private" {
		_, err := bot.SendMessage(u.Message.Chat.Id, texts.ThisChatIsPrivateICanNotToUnpinMessage)
		return err
	}
	_, err = bot.UnpinChatMessage(u.Message.Chat.Id)
	return err
}

func Promote(bot ext.Bot, u *gotgbot.Update, args []string) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		return nil
	}
	if !utils.MemberCanPromote(bot, u) {
		return nil
	}
	if !utils.BotCanPromote(bot, u) {
		return nil
	}
	var promoteId int
	if u.Message.ReplyToMessage != nil {
		promoteId = u.Message.ReplyToMessage.From.Id
	} else {
		if len(args) > 0 {
			promoteId, err = strconv.Atoi(args[0])
		} else {
			_, err = bot.SendMessage(u.Message.Chat.Id, texts.PleaseReplyToTheMessageOfThePersonYouWantToGrantAdministratorRightsToOrEnterTheirID)
			return err
		}
	}
	_, err = bot.PromoteChatMember(u.Message.Chat.Id, promoteId)
	return  err
}

func Demote(bot ext.Bot, u *gotgbot.Update, args []string) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		return nil
	}
	if !utils.MemberCanPromote(bot, u) {
		return nil
	}
	if !utils.BotCanPromote(bot, u) {
		return nil
	}
	var promoteId int
	if u.Message.ReplyToMessage != nil {
		promoteId = u.Message.ReplyToMessage.From.Id
	} else {
		if len(args) > 0 {
			promoteId, err = strconv.Atoi(args[0])
		} else {
			_, err = bot.SendMessage(u.Message.Chat.Id, texts.PleaseReplyToTheMessageOfThePersonYouWantToRemoveAdministratorRightsToOrEnterTheirID)
			return err
		}
	}
	_, err = bot.DemoteChatMember(u.Message.Chat.Id, promoteId)
	return  err
}

func Purge(bot ext.Bot, u *gotgbot.Update) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		return nil
	}
	if !utils.MemberCanDelMsg(bot, u) {
		return nil
	}
	if !utils.BotCanDelMsg(bot, u) {
		return nil
	}
	chatId := u.Message.Chat.Id
	lastId := u.Message.ReplyToMessage.MessageId
	for i := u.Message.MessageId; i >= lastId; i-- {
		_, err := bot.DeleteMessage(chatId, i)
		errors.SendError(err)
	}
	msg, err := bot.SendMessage(chatId, texts.PurgeCompleted)
	time.Sleep(5000)
	_, err = bot.DeleteMessage(chatId, msg.MessageId)
	return nil
}

func DisableCommands(bot ext.Bot, u *gotgbot.Update, args []string) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		return nil
	}
	if len(args) == 0 {
		text := "Эта команда отключает пользовательские команды. Список команд:"
		for _, command := range strings.Split(userCommands, " ") {
			text += "\n• <code>" + command + "</code>"
		}
		_, err = bot.SendMessageHTML(u.Message.Chat.Id, text)
		return err
	}
	disabledCommands := ""
	for _, command := range args {
		if strings.Contains(userCommands, command) {
			disabledCommands += command
		}
	}
	if disabledCommands == "" {
		_, err = bot.SendMessage(u.Message.Chat.Id, "Вы не написали ни одной пользовательской команды.")
		return err
	}
	err = sql.SetDisabledCommands(u.Message.Chat.Id, disabledCommands)
	if err != nil {
		return err
	}
	_, err = bot.SendMessage(u.Message.Chat.Id, "Отключены следующие команды: " + disabledCommands)
	return err
}

func GetDisabledCommands(bot ext.Bot, u *gotgbot.Update) error {
	disabledCommands, err := sql.GetDisabledCommands(u.Message.Chat.Id)
	if err != nil {
		return err
	}
	if disabledCommands == "" {
		_, err = bot.SendMessage(u.Message.Chat.Id, "Все команды включены.")
	} else {
		_, err = bot.SendMessage(u.Message.Chat.Id, "Отключены следующие команды: " + disabledCommands)
	}
	return err
}

func DisableAllCommands(bot ext.Bot, u *gotgbot.Update) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		return nil
	}
	err = sql.SetDisabledCommands(u.Message.Chat.Id, userCommands)
	if err != nil {
		return err
	}
	_, err = bot.SendMessage(u.Message.Chat.Id, "Все команды отключены.")
	return err
}

func EnableAllCommands(bot ext.Bot, u *gotgbot.Update) error {
	userMember, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(userMember) {
		return nil
	}
	err = sql.SetDisabledCommands(u.Message.Chat.Id, "")
	if err != nil {
		return err
	}
	_, err = bot.SendMessage(u.Message.Chat.Id, "Все команды включены.")
	return err
}