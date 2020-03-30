package rules

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/utils"
	"strings"
)

func SetRules(b ext.Bot, u *gotgbot.Update, args []string) error {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	if len(args) == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, "Эта комманда позволяет установить правила.")
		return err
	}
	var rules string
	for _, value := range args {
		rules += value + " "
	}
	rules = strings.TrimSuffix(rules, " ")
	err = sql.SetRules(u.Message.Chat.Id, rules)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(u.Message.Chat.Id, "Правила установлены! Вы можете посмотреть их с помощью команды /rules.")
	return err
}

func GetRules(b ext.Bot, u *gotgbot.Update) error {
	rules, err := sql.GetRules(u.Message.Chat.Id)
	if err != nil {
		errors.SendError(err)
		_, err = b.SendMessage(u.Message.Chat.Id, "Ошибка получения правил.\n" + err.Error())
		return err;
	}
	_, err = b.SendMessage(u.Message.Chat.Id, rules)
	return err;
}