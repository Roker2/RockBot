package bans

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/sqlite"
	"github.com/Roker2/RockBot/modules/utils"
	"log"
	"strconv"
)

func Ban(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	banId, errortext := utils.ExtractId(b, u, args)
	if banId == 0 {
		_, err := b.SendMessage(chat.Id, errortext)
		return err
	}
	if utils.ItIsMe(b, u, banId) {
		return nil
	}
	log.Print(strconv.Itoa(banId))
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	banMember, err := chat.GetMember(banId)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	/*if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не не имеете права что-то делать с пользователями.")
		return err
	}*/
	utils.MemberCanRestrictMembers(b, u)
	if banMember != nil {
		if utils.MemberIsAdministrator(banMember) {
			_, err = b.SendMessage(u.Message.Chat.Id, "Я не могу забанить администратора.")
			return err
		}
	}
	_, err = chat.KickMember(banId)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " забанен!")
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
	log.Print(strconv.Itoa(banId))
	if utils.IsUserInChat(chat, banId) {
		_, err := b.SendMessage(chat.Id, "Этот пользователь в данный момент в чате.")
		return err
	}
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	banMember, err := chat.GetMember(banId)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(chat.Id, "Вы не администратор.")
		return err
	}
	/*if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не не имеете права что-то делать с пользователями.")
		return err
	}*/
	utils.MemberCanRestrictMembers(b, u)
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(chat.Id, "Я не могу сделать это с администратором.")
		return err
	} else {
		_, err = chat.UnbanMember(banId)
		if err != nil {
			return err
		}
		sqlite.ResetUserWarns(u.Message.Chat.Id, banId)
		_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " разбанен!")
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
	log.Print(strconv.Itoa(banId))
	if !utils.IsUserInChat(chat, banId) {
		_, err := b.SendMessage(chat.Id, "Этого пользователя нет в чате.")
		return err
	}
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	banMember, err := chat.GetMember(banId)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(chat.Id, "Вы не администратор.")
		return err
	}
	/*if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не не имеете права что-то делать с пользователями.")
		return err
	}*/
	utils.MemberCanRestrictMembers(b, u)
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(chat.Id, "Я не могу сделать это с администратором.")
		return err
	} else {
		_, err = chat.UnbanMember(banId)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " кикнут!")
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
		_, err = b.SendMessage(chat.Id, "Я не могу сделать это с администратором.")
		return err
	}
	_, err = chat.UnbanMember(banId)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " кикнут!")
	return err
}
