package info

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"strconv"
)

func UserInfo(b ext.Bot, u *gotgbot.Update) error {
	user := u.Message.From
	if u.Message.ReplyToMessage != nil {
		user = u.Message.ReplyToMessage.From
	}
	textInfo := "<b>First Name:</b> " + user.FirstName
	if len(user.LastName) != 0 {
		textInfo += "\n<b>Last Name:</b> " + user.LastName
	}
	textInfo += "\n<b>User ID:</b> <code>" + strconv.Itoa(user.Id)
	if len(user.Username) != 0 {
		textInfo += "</code>\n<b>User Name:</b> @" + user.Username
	}
	textInfo += "\n<b>Bot:</b> " + strconv.FormatBool(user.IsBot)
	_, err := b.SendMessageHTML(u.Message.Chat.Id, textInfo)
	return err
}

func ChatInfo(b ext.Bot, u *gotgbot.Update) error {
	chat := u.Message.Chat
	membersCount, err := chat.GetMembersCount()
	if err != nil {
		return err
	}
	textInfo := "<b>Chat ID:</b> <code>" + strconv.Itoa(chat.Id) + "</code>\n<b>User Name:</b> @" + chat.Username + "\n<b>Members count:</b> " + strconv.Itoa(membersCount)
	_, err = b.SendMessageHTML(u.Message.Chat.Id, textInfo)
	return err
}