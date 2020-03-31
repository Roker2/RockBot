package welcome

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/utils"
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
			_, err = b.SendMessageHTML(u.Message.Chat.Id, strings.ReplaceAll(strings.ReplaceAll(welcome, "{firstName}", member.FirstName), "<br>", "\n"))
		} else {
			_, err = b.SendMessageHTML(u.Message.Chat.Id, strings.ReplaceAll(strings.ReplaceAll(welcome, "{firstName}", "пользователь"), "<br>", "\n"))
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

func SetWelcome(b ext.Bot, u *gotgbot.Update, args []string) error {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	if len(args) == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, "Эта комманда позволяет заменить встречающую реплику на свою.\nФорматирование:\n{firstName} - Имя пользователя\n{имя переменной} заменяется на текстовое значение.\nИспользуйте HTML для форматирования текста. <br> - переход на новую строку.")
		return err
	}
	welcome := u.Message.Text
	welcome = strings.ReplaceAll(welcome, "/setwelcome ", "")
	err = sql.SetWelcome(u.Message.Chat.Id, welcome)
	return err
}