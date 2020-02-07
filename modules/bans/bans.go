package bans

import (
	"../utils"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"log"
	"strconv"
)

func Ban(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	var banId int
	if len(args) >= 1  {
		banId2, err := strconv.Atoi(args[0])
		if err != nil {
			_, err = b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста ID, а не бред.")
			if err != nil {
				return err
			}
		}
		banId = banId2
	} else {
		if !utils.IsReply(b, u) {
			return nil
		}
		banId = u.Message.ReplyToMessage.From.Id
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
	var banId int
	if len(args) >= 1  {
		banId2, err := strconv.Atoi(args[0])
		if err != nil {
			_, err = b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста ID, а не бред.")
			if err != nil {
				return err
			}
		}
		banId = banId2
	} else {
		if !utils.IsReply(b, u) {
			return nil
		}
		banId = u.Message.ReplyToMessage.From.Id
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
		_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " разбанен!")
	}
	return err
}

func Kick(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	var banId int
	if len(args) >= 1  {
		banId2, err := strconv.Atoi(args[0])
		if err != nil {
			_, err = b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста ID, а не бред.")
			if err != nil {
				return err
			}
		}
		banId = banId2
	} else {
		if !utils.IsReply(b, u) {
			return nil
		}
		banId = u.Message.ReplyToMessage.From.Id
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
