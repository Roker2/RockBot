package welcome

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"strings"
)

func NewMember(b ext.Bot, u *gotgbot.Update) error {
	newMembers := u.EffectiveMessage.NewChatMembers
	for _, member := range newMembers {
		if member.Id == b.Id {
			continue
		}
		welcome, err := sql.GetWelcome(u.Message.Chat.Id)
		if err != nil {
			errors.SendError(err)
		}
		if member.FirstName != "" {
			_, err = b.SendMessage(u.Message.Chat.Id, strings.ReplaceAll(welcome, "{firstName}", member.FirstName))
		} else {
			_, err = b.SendMessage(u.Message.Chat.Id, strings.ReplaceAll(welcome, "{firstName}", "пользователь"))
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