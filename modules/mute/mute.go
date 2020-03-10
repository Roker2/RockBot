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
	muteId := utils.ExtractId(b, u, args)
	if muteId == 0 {
		_, err := b.SendMessage(chat.Id, "Что-то не так, не могу получить ID пользователя.")
		return err
	}
	if muteId == -1 {
		_, err := b.SendMessage(chat.Id, "Я для тебя что, вещь?")
		return err
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
	muteId := utils.ExtractId(b, u, args)
	if muteId == 0 {
		_, err := b.SendMessage(chat.Id, "Что-то не так, не могу получить ID пользователя.")
		return err
	}
	if muteId == -1 {
		_, err := b.SendMessage(chat.Id, "Это бессмысленно.")
		return err
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