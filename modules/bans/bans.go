package bans

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/utils"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/sirupsen/logrus"
	"strconv"
)

func commonBan(b ext.Bot, u *gotgbot.Update, args []string) (bool, int, error) {
	banId, errorText := utils.ExtractId(b, u, args)
	if banId == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, errorText)
		return false, 0, err
	}
	if utils.ItIsMe(b, u, banId) {
		return false, 0, nil
	}
	if !utils.BotIsAdministrator(b, u) {
		return false, 0, nil
	}
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return false, 0, err
	}
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return false, 0, err
	}
	if !utils.MemberCanRestrictMembers(b, u) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotToDoSomethingWithUsers)
		return false, 0, err
	}
	return true, banId, nil
}

func Ban(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, banId, err := commonBan(b, u, args)
	if !canBan {
		return err
	}
	banMember, err := u.Message.Chat.GetMember(banId)
	if err != nil {
		return err
	}
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotToBanAdministrator)
		return err
	}
	_, err = u.Message.Chat.KickMember(banId)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsBanned(banMember.User.FirstName))
	return err
}

func Unban(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	banId, errortext := utils.ExtractId(b, u, args)
	if banId == 0 {
		_, err := b.SendMessage(chat.Id, errortext)
		return err
	}
	if utils.ItIsMe(b, u, banId) {
		return nil
	}
	logrus.Println(strconv.Itoa(banId))
	if utils.IsUserInChat(chat, banId) {
		_, err := b.SendMessage(chat.Id, texts.UserIsInTheChat)
		return err
	}
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	banMember, err := chat.GetMember(banId)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotToDoSomethingWithUsers)
		return err
	}
	utils.MemberCanRestrictMembers(b, u)
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(chat.Id, texts.ICanNotToBanAdministrator)
		return err
	} else {
		_, err = chat.UnbanMember(banId)
		if err != nil {
			return err
		}
		err = sql.ResetUserWarns(u.Message.Chat.Id, banId)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(chat.Id, texts.UserIsUnbanned(banMember.User.FirstName))
	}
	return err
}

func Kick(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	banId, errortext := utils.ExtractId(b, u, args)
	if banId == 0 {
		_, err := b.SendMessage(chat.Id, errortext)
		return err
	}
	if utils.ItIsMe(b, u, banId) {
		return nil
	}
	logrus.Println(strconv.Itoa(banId))
	if !utils.IsUserInChat(chat, banId) {
		_, err := b.SendMessage(chat.Id, texts.UserIsNotInTheChat)
		return err
	}
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	banMember, err := chat.GetMember(banId)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouCanNotToDoSomethingWithUsers)
		return err
	}
	utils.MemberCanRestrictMembers(b, u)
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(chat.Id, texts.ICanNotToKickAdministrator)
		return err
	}
	if !utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(chat.Id, texts.YouAreNotAdministrator)
		return err
	} else {
		_, err = chat.UnbanMember(banId)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(chat.Id, texts.UserIsKicked(banMember.User.FirstName))
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
		_, err = b.SendMessage(chat.Id, texts.ICanNotToKickAdministrator)
		return err
	}
	_, err = chat.UnbanMember(banId)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(chat.Id, texts.UserIsKicked(banMember.User.FirstName))
	return err
}
