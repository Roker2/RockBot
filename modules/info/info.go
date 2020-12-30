package info

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
	"strconv"
)

func UserInfo(b ext.Bot, u *gotgbot.Update, args []string) error {
	userId, errorText := utils.ExtractId(b, u, args)
	if userId == 0 {
		if errorText == texts.ReplyOrWriteId {
			userId = u.Message.From.Id
		} else {
			_, err := b.SendMessage(u.Message.Chat.Id, errorText)
			return err
		}
	}
	chatMember, err := u.Message.Chat.GetMember(userId)
	if err != nil {
		return err
	}
	textInfo := "<b>First Name:</b> " + chatMember.User.FirstName
	if len(chatMember.User.LastName) != 0 {
		textInfo += "\n<b>Last Name:</b> " + chatMember.User.LastName
	}
	textInfo += "\n<b>User ID:</b> <code>" + strconv.Itoa(chatMember.User.Id)
	if len(chatMember.User.Username) != 0 {
		textInfo += "</code>\n<b>User Name:</b> @" + chatMember.User.Username
	}
	textInfo += "\n<b>Bot:</b> " + strconv.FormatBool(chatMember.User.IsBot)
	_, err = b.SendMessageHTML(u.Message.Chat.Id, textInfo)
	return err
}

func ChatInfo(b ext.Bot, u *gotgbot.Update) error {
	chat := u.Message.Chat
	membersCount, err := chat.GetMembersCount()
	if err != nil {
		return err
	}
	textInfo := "<b>Chat ID:</b> <code>" + strconv.Itoa(chat.Id) + "</code>"
	if len(chat.Username) != 0 {
		textInfo += "\n<b>User Name:</b> @" + chat.Username
	}
	textInfo += "\n<b>Members count:</b> " + strconv.Itoa(membersCount)
	_, err = b.SendMessageHTML(u.Message.Chat.Id, textInfo)
	return err
}

func SaveUserToDatabase(b ext.Bot, u *gotgbot.Update) error {
	err := sql.SaveUser(u.Message.From)
	if err != nil {
		return err
	}
	if u.Message.ForwardFrom != nil {
		err = sql.SaveUser(u.Message.ForwardFrom)
		if err != nil {
			return err
		}
	}
	if utils.IsReply(b, u, false) {
		err = sql.SaveUser(u.Message.ReplyToMessage.From)
		if err != nil {
			return err
		}
	}
	return err
}