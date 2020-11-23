package bans

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
)

func Ban(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, banId, err := utils.CommonBan(b, u, args)
	if !canBan || err != nil {
		return err
	}
	banMember, err := u.Message.Chat.GetMember(banId)
	if err != nil {
		return err
	}
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotBanAdministrator)
		return err
	}
	_, err = u.Message.Chat.KickMember(banId)
	if err != nil {
		return err
	}
	if banMember.User.Username == "" {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsBanned(banMember.User.FirstName))
	} else {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsBanned("@" + banMember.User.Username))
	}
	return err
}

func Unban(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, banId, err := utils.CommonBan(b, u, args)
	if !canBan || err != nil {
		return err
	}
	banMember, err := u.Message.Chat.GetMember(banId)
	if err != nil {
		return err
	}
	if utils.IsUserInChat(u.Message.Chat, banId) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsInTheChat)
		return err
	}
	_, err = u.Message.Chat.UnbanMember(banId)
	if err != nil {
		return err
	}
	err = sql.ResetUserWarns(u.Message.Chat.Id, banId)
	if err != nil {
		return err
	}
	if banMember.User.Username == "" {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsUnbanned(banMember.User.FirstName))
	} else {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsUnbanned("@" + banMember.User.Username))
	}
	return err
}

func Kick(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, banId, err := utils.CommonBan(b, u, args)
	if !canBan || err != nil {
		return err
	}
	banMember, err := u.Message.Chat.GetMember(banId)
	if err != nil {
		return err
	}
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotKickAdministrator)
		return err
	}
	_, err = u.Message.Chat.UnbanMember(banId)
	if err != nil {
		return err
	}
	if banMember.User.Username == "" {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsKicked(banMember.User.FirstName))
	} else {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsKicked("@" + banMember.User.Username))
	}
	return err
}

func Kickme(b ext.Bot, u *gotgbot.Update) error {
	chat := u.Message.Chat
	banId := u.Message.From.Id
	if !utils.BotIsAdministrator(b, u) {
		return nil
	}
	banMember, err := chat.GetMember(banId)
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(chat.Id, texts.ICanNotKickAdministrator)
		return err
	}
	_, err = chat.UnbanMember(banId)
	if err != nil {
		return err
	}
	if banMember.User.Username == "" {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsKicked(banMember.User.FirstName))
	} else {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsKicked("@" + banMember.User.Username))
	}
	return err
}
