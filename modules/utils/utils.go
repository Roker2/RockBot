package utils

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"strconv"
)

func MemberIsCreator(member *ext.ChatMember) bool {
	return member.Status == "creator"
}

func MemberIsAdministrator(member *ext.ChatMember) bool {
	return member.Status == "administrator" || member.Status == "creator"
}

func IsReply(b ext.Bot, u *gotgbot.Update, writeMsg bool) bool {
	if u.Message.ReplyToMessage == nil {
		if writeMsg {
			_, err := b.SendMessage(u.Message.Chat.Id, texts.ReplyOrWriteId)
			errors.SendError(err)
		}
		return false
	}
	return true
}

func IsForward(msg *ext.Message) bool {
	return msg.ForwardFrom != nil
}

func BotIsAdministrator(b ext.Bot, u *gotgbot.Update) bool {
	botmember, err := u.Message.Chat.GetMember(b.Id)
	if !MemberIsAdministrator(botmember) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.IAmNotAdministrator)
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
		_, err := b.SendMessage(u.Message.Chat.Id, texts.AmIThing)
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
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotPinOrUnpin)
		errors.SendError(err)
		return false
	}
	return true
}

func BotCanPin(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(b.Id)
	errors.SendError(err)
	if !member.CanPinMessages {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotPinOrUnpin)
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
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotPromote)
		errors.SendError(err)
		return false
	}
	return true
}

func BotCanPromote(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(b.Id)
	errors.SendError(err)
	if !member.CanPromoteMembers {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotPromote)
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
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotDeleteMessages)
		errors.SendError(err)
		return false
	}
	return true
}

func BotCanDelMsg(b ext.Bot, u *gotgbot.Update) bool {
	member, err := u.Message.Chat.GetMember(b.Id)
	errors.SendError(err)
	if !member.CanDeleteMessages {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotDeleteMessages)
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
	if !member.CanRestrictMembers && !MemberIsCreator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotDoSomethingWithUsers)
		errors.SendError(err)
		return false
	}
	return true
}

func ExtractId(b ext.Bot, u *gotgbot.Update, args []string) (int, string) {
	if IsForward(u.Message) {
		return u.Message.ForwardFrom.Id, ""
	}
	if IsReply(b, u, false) {
		if IsForward(u.Message.ReplyToMessage) {
			return u.Message.ReplyToMessage.ForwardFrom.Id, ""
		} else {
			return u.Message.ReplyToMessage.From.Id, ""
		}
	} else {
		id := 0
		if len(args) >= 1  {
			if args[0][0] == '@' {
				var err error
				id, err = sql.GetUserId(args[0][1:])
				if err != nil {
					errors.SendError(err)
					return id, texts.ICanNotGetId
				}
				return id, ""
			} else {
				id, err := strconv.Atoi(args[0])
				if err != nil {
					errors.SendError(err)
					return 0, texts.WriteIdNotBadText
				}
				return id, ""
			}
		} else {
			return 0, texts.ReplyOrWriteId
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
	return msg[(index + 1):]
}

func CommonBan(b ext.Bot, u *gotgbot.Update, args []string) (bool, int, error) {
	banId, errorText := ExtractId(b, u, args)
	if banId == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, errorText)
		return false, 0, err
	}
	if ItIsMe(b, u, banId) {
		return false, 0, nil
	}
	if !BotIsAdministrator(b, u) {
		return false, 0, nil
	}
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return false, 0, err
	}
	if !MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return false, 0, err
	}
	if !MemberCanRestrictMembers(b, u) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotDoSomethingWithUsers)
		return false, 0, err
	}
	return true, banId, nil
}