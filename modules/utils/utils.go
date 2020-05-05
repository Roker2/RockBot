package utils

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"log"
	"strconv"
)

func MemberIsCreator(member *ext.ChatMember) bool {
	if member.Status == "creator" {
		return true
	}
	return false
}

func MemberIsAdministrator(member *ext.ChatMember) bool {
	if member.Status == "administrator" || member.Status == "creator" {
		return true
	}
	return false
}

func IsReply(b ext.Bot, u *gotgbot.Update, writeMsg bool) bool {
	if u.Message.ReplyToMessage == nil {
		if writeMsg {
			_, err := b.SendMessage(u.Message.Chat.Id, "Ответьте пожалуйста на сообщение того, с кем Вы хотите что-то сделать.")
			errors.SendError(err)
		}
		return false
	}
	return true
}

func BotIsAdministrator(b ext.Bot, u *gotgbot.Update) bool {
	botmember, err := u.Message.Chat.GetMember(b.Id)
	if !MemberIsAdministrator(botmember) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Я не администратор.")
		errors.SendError(err)
	}
	return MemberIsAdministrator(botmember)
}

func IsUserInChat(chat *ext.Chat, userId int) bool  {
	member, err := chat.GetMember(userId)
	errors.SendError(err)
	if member.Status == "left" || member.Status == "kicked" {
		return false
	}
	return true
}

func ItIsMe(b ext.Bot, u *gotgbot.Update, Id int) bool {
	if b.Id == Id {
		_, err := b.SendMessage(u.Message.Chat.Id, "Я для тебя что, вещь, которую можно выбросить?")
		errors.SendError(err)
		return true
	}
	return false
}

func MemberCanPin(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	errors.SendError(err)
	if MemberIsCreator(member) == true {
		return true
	}
	if !member.CanPinMessages {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не имеете права закреплять или откреплять сообщения.")
		errors.SendError(err)
		return false
	}
	return true
}

func BotCanPin(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(b.Id)
	errors.SendError(err)
	if !member.CanPinMessages {
		_, err = b.SendMessage(u.Message.Chat.Id, "Я не имею права закреплять или откреплять сообщения.")
		errors.SendError(err)
		return false
	}
	return true
}

func MemberCanPromote(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	errors.SendError(err)
	if MemberIsCreator(member) == true {
		return true
	}
	if !member.CanPromoteMembers {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не имеете права выдавать права администратора.")
		errors.SendError(err)
		return false
	}
	return true
}

func BotCanPromote(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(b.Id)
	errors.SendError(err)
	if !member.CanPromoteMembers {
		_, err = b.SendMessage(u.Message.Chat.Id, "Я не имею права выдавать права администратора.")
		errors.SendError(err)
		return false
	}
	return true
}

func MemberCanDelMsg(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	errors.SendError(err)
	if MemberIsCreator(member) == true {
		return true
	}
	if !member.CanDeleteMessages {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не имеете права удалять сообщения.")
		errors.SendError(err)
		return false
	}
	return true
}

func BotCanDelMsg(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(b.Id)
	errors.SendError(err)
	if !member.CanDeleteMessages {
		_, err = b.SendMessage(u.Message.Chat.Id, "Я не имею права удалять сообщения.")
		errors.SendError(err)
		return false
	}
	return true
}

func MemberCanRestrictMembers(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	errors.SendError(err)
	if MemberIsCreator(member) == true {
		return true
	}
	if !member.CanRestrictMembers {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не не имеете права что-то делать с пользователями.")
		errors.SendError(err)
		return false
	}
	return true
}

func ExtractId(b ext.Bot, u *gotgbot.Update, args []string) (int, string) {
	if IsReply(b, u, false) {
		return u.Message.ReplyToMessage.From.Id, ""
	} else {
		if len(args) >= 1  {
			banId2, err := strconv.Atoi(args[0])
			if err != nil {
				return 0, "Введите пожалуйста ID, а не бред."
			}
			return banId2, ""
		} else {
			return 0, "Ответьте пожалуйста на сообщение человека, с которым Вы хотите что-то сделать, или напишите его ID"
		}
	}
}

func RemoveCommand(msg string) string {
	var index int
	for valueIndex, value := range msg {
		if (value == ' ') || (value == '\n') {
			index = valueIndex
			break
		}
	}
	log.Print(msg)
	return msg[(index + 1):]
}