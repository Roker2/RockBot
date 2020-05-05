package mute

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/utils"
	"log"
	"strconv"
	"time"
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

func TemporarlyMute(b ext.Bot, u *gotgbot.Update, args []string) error {
	chat := u.Message.Chat
	if len(args) == 0 {
		_, err := b.SendMessage(chat.Id, "Данная команда предназначена для временного mute.")
		return err
	}
	timeInterval := time.Now().Unix()
	for _, temp := range args {
		tempTime, err := strconv.Atoi(temp[:len(temp) - 1])
		if err != nil {
			errors.SendError(err)
			continue
		}
		switch temp[len(temp) - 1] {
		case 'm':
			timeInterval += int64(tempTime * 60)
		case 'h':
			timeInterval += int64(tempTime * 60 * 60)
		case 'd':
			timeInterval += int64(tempTime * 60 * 60 * 24)
		}
	}
	muteId, errortext := utils.ExtractId(b, u, args)
	if muteId == 0 {
		if !utils.IsReply(b, u, false) {
			_, err := b.SendMessage(chat.Id, errortext)
			return err
		} else {
			muteId = u.Message.ReplyToMessage.From.Id
		}
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
		newRestrict := b.NewSendableRestrictChatMember(chat.Id, muteId)
		newRestrict.UntilDate = timeInterval
		_, err = newRestrict.Send()
		if err != nil {
			return err
		}
		_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " теперь будет молчать.")
	}
	return err
}