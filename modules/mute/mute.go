package mute

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/utils"
	"log"
	"strconv"
)

func Mute(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	muteId, errortext := utils.ExtractId(b, u, args)
	if muteId == 0 {
		_, err := b.SendMessage(chat.Id, errortext)
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
	if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
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
	muteId, errortext := utils.ExtractId(b, u, args)
	if muteId == 0 {
		_, err := b.SendMessage(chat.Id, errortext)
		return err
	}
	log.Print(strconv.Itoa(muteId))
	member, err := chat.GetMember(u.Message.From.Id)
	if !utils.BotIsAdministrator(b, u) {
		return err
	}
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	if !member.CanRestrictMembers && !utils.MemberIsCreator(member) {
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