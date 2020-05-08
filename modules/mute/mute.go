package mute

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
	"strconv"
	"time"
)

func Mute(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, muteId, err := utils.CommonBan(b, u, args)
	if !canBan {
		return err
	}
	muteMember, err := u.Message.Chat.GetMember(muteId)
	if err != nil {
		return err
	}
	if utils.MemberIsAdministrator(muteMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotToMuteAdministrator)
		return err
	} else {
		_, err = b.RestrictChatMember(u.Message.Chat.Id, muteId)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsMuted(muteMember.User.FirstName))
	}
	return err
}

func Unmute(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, muteId, err := utils.CommonBan(b, u, args)
	if !canBan {
		return err
	}
	_, err = b.UnRestrictChatMember(u.Message.Chat.Id, muteId)
	if err != nil {
		return err
	}
	_, err = b.SendMessage(u.Message.Chat.Id, texts.TalkToMeHere)
	return err
}

func TemporaryMute(b ext.Bot, u *gotgbot.Update, args []string) error {
	if len(args) == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, texts.ThisCommandForTemporaryMute)
		return err
	}
	canBan, muteId, err := utils.CommonBan(b, u, args)
	if !canBan {
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
	muteMember, err := u.Message.Chat.GetMember(muteId)
	if err != nil {
		return err
	}
	if utils.MemberIsAdministrator(muteMember) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.ICanNotToMuteAdministrator)
		return err
	}
	newRestrict := b.NewSendableRestrictChatMember(u.Message.Chat.Id, muteId)
	newRestrict.UntilDate = timeInterval
	_, err = newRestrict.Send()
	if err != nil {
		return err
	}
	_, err = b.SendMessage(u.Message.Chat.Id, texts.UserIsMuted(muteMember.User.FirstName))
	return err
}