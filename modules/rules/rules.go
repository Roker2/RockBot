package rules

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
	"strings"
)

func SetRules(b ext.Bot, u *gotgbot.Update, args []string) error {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	if len(args) == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, texts.DescriptionOfRulesCommand)
		return err
	}
	rules := utils.RemoveCommand(u.Message.OriginalHTML())
	err = sql.SetRules(u.Message.Chat.Id, rules)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(u.Message.Chat.Id, texts.NewRulesWereAdded)
	return err
}

func GetRules(b ext.Bot, u *gotgbot.Update) error {
	rules, err := sql.GetRules(u.Message.Chat.Id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		errors.SendError(err)
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ErrorOfGettingRules(err.Error()))
		return err;
	}
	_, err = b.SendMessageHTML(u.Message.Chat.Id, strings.ReplaceAll(rules, "<br>", "\n"))
	return err;
}