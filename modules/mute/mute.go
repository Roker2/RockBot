package mute

import (
	"../utils"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"log"
	"strconv"
)

func Mute(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	var muteId int
	if len(args) >= 1  {
		muteId2, err := strconv.Atoi(args[0])
		if err != nil {
			_, _ = b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста ID, а не бред.")
			return err
		}
		muteId = muteId2
	} else {
		if !utils.IsReply(b, u) {
			return nil
		}
		muteId = u.Message.ReplyToMessage.From.Id
	}
	if utils.ItIsMe(b, u, muteId) {
		return nil
	}
	log.Print(strconv.Itoa(muteId))
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	banMember, err := chat.GetMember(muteId)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	if !member.CanRestrictMembers {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не не имеете права что-то делать с пользователями.")
		return err
	}
	if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Я не могу заставить молчать администратора.")
		return err
	} else {
		_, err = b.RestrictChatMember(chat.Id, muteId)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " теперь будет молчать.")
	}
	return err
}

func Unmute(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	var muteId int
	if len(args) >= 1  {
		muteId2, err := strconv.Atoi(args[0])
		if err != nil {
			_, _ = b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста ID, а не бред.")
			return err
		}
		muteId = muteId2
	} else {
		if !utils.IsReply(b, u) {
			return nil
		}
		muteId = u.Message.ReplyToMessage.From.Id
	}
	if utils.ItIsMe(b, u, muteId) {
		return nil
	}
	log.Print(strconv.Itoa(muteId))
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	//banMember, err := chat.GetMember(muteId)
	/*if utils.MemberIsAdministrator(banMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Я не могу заставить молчать администратора.")
		return err
	}*/
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	if !member.CanRestrictMembers {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не не имеете права что-то делать с пользователями.")
	} else {
		_, err = b.UnRestrictChatMember(chat.Id, muteId)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(chat.Id, "Поговори мне тут...")
	}
	return err
}