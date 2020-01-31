package welcome

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
)

func NewMember(b ext.Bot, u *gotgbot.Update) error {
	newMembers := u.EffectiveMessage.NewChatMembers
	for _, member := range newMembers {
		if member.Id == b.Id {
			continue
		}
		var err error
		if member.FirstName != "" {
			_, err = b.SendMessage(u.Message.Chat.Id, "Добро пожаловать, " + member.FirstName + "!")
		} else {
			_, err = b.SendMessage(u.Message.Chat.Id, "Добро пожаловать, пользователь!")
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func LeftMember(b ext.Bot, u *gotgbot.Update) error {
	member := u.EffectiveMessage.LeftChatMember
	var err error
	if member.FirstName != "" {
		_, err = b.SendMessage(u.Message.Chat.Id, "До встречи, " + member.FirstName + "!")
	} else {
		_, err = b.SendMessage(u.Message.Chat.Id, "До встречи, пользователь!")
	}
	if err != nil {
		return err
	}
	return nil
}