package warns
//version: 1,0

import (
	"../sqlite"
	"../utils"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"log"
	"strconv"
)

func WarnUser(b ext.Bot, u *gotgbot.Update, args []string) error {
	if !utils.BotIsAdministrator(b, u) {
		return nil
	}
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
 	if !utils.MemberIsAdministrator(member) {
 		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
 		return err
 	}
	log.Print(strconv.Itoa(banId))
  	utils.MemberCanRestrictMembers(b, u)
 	banMember, err := chat.GetMember(banId)
	if banMember != nil {
		if utils.MemberIsAdministrator(banMember) {
			_, err = b.SendMessage(u.Message.Chat.Id, "Я не могу дать предупреждение администратору.")
			return err
		}
	}
  	quantity := sqlite.AddUserWarn(u.Message.Chat.Id, banId)
  	if quantity == -1 {
    	_, err = b.SendMessage(u.Message.Chat.Id, "У нас маленько проблема... @Roker2!")
    	if err != nil {
      	return err
    	}
  	} else {
  		_, err := b.SendMessage(u.Message.Chat.Id, "Количество предупреждений у " + banMember.User.FirstName + ": " + strconv.Itoa(quantity) + "/" + strconv.Itoa(sqlite.GetWarnsQuantityOfChat(u.Message.Chat.Id)))
    	if err != nil {
      	return err
   	 	}
  	}
  	if strconv.Itoa(quantity) >= strconv.Itoa(sqlite.GetWarnsQuantityOfChat(u.Message.Chat.Id)) {
  		_, err = chat.KickMember(banId)
  		if err != nil {
  			return err
  		}
  		_, err = b.SendMessage(chat.Id, "Пользователь " + banMember.User.FirstName + " забанен!")
  		return err
  	}
  	return nil
}

func WarnsQuantity (b ext.Bot, u *gotgbot.Update) error {
  	_, err := b.SendMessage(u.Message.Chat.Id, strconv.Itoa(sqlite.GetWarnsQuantityOfChat(u.Message.Chat.Id)))
  	return err
}

func GetUserWarns(b ext.Bot, u *gotgbot.Update, args []string) error {
	banId, errortext := utils.ExtractId(b, u, args)
	if banId == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, errortext)
		return err
	}
 	if utils.ItIsMe(b, u, banId) {
 		return nil
 	}
 	log.Print(strconv.Itoa(banId))
  	if sqlite.GetUserWarns(u.Message.Chat.Id, banId) == -1 {
    	_, err := b.SendMessage(u.Message.Chat.Id, "У нас маленько проблема... @Roker2!")
    	if err != nil {
      return err
    }
  	} else {
 	  	banMember, err := u.Message.Chat.GetMember(banId)
    	if err != nil {
      return err
    }
    	_, err = b.SendMessage(u.Message.Chat.Id, "Количество предупреждений у " + banMember.User.FirstName + ": " + strconv.Itoa(sqlite.GetUserWarns(u.Message.Chat.Id, banId)) + "/" + strconv.Itoa(sqlite.GetWarnsQuantityOfChat(u.Message.Chat.Id)))
    	if err != nil {
      return err
    }
  	}
  	return nil
}

func SetWarnsQuantity (b ext.Bot, u *gotgbot.Update, args []string) error {
	if !utils.BotIsAdministrator(b, u) {
		return nil
	}
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if err != nil {
		return err
	}
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
		return err
	}
	//var quantity int
  	if len(args) >= 1 {
    	quantity, err := strconv.Atoi(args[0])
    	if err != nil {
    		_, err = b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста число, а не бред.")
 			if err != nil {
 				return err
 			}
    }
    sqlite.SetWarnsQuantityOfChat(u.Message.Chat.Id, quantity)
  	} else {
    	_, err := b.SendMessage(u.Message.Chat.Id, "Введите пожалуйста число.")
    	if err != nil {
      return err
    }
  	}
  	return nil
}

func ResetWarns (b ext.Bot, u *gotgbot.Update, args []string) error {
	if !utils.BotIsAdministrator(b, u) {
		return nil
	}
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
    	if !utils.IsReply(b, u, true) {
      		return nil
    	}
    	banId = u.Message.ReplyToMessage.From.Id
  	}
  	if utils.ItIsMe(b, u, banId) {
    	return nil
  	}
  	log.Print(strconv.Itoa(banId))
  	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
  	if !utils.BotIsAdministrator(b, u) {
    	return err
  	}
  	if !utils.MemberIsAdministrator(member) {
    	_, err = b.SendMessage(u.Message.Chat.Id, "Вы не администратор.")
    	return err
  	}
  	log.Print(strconv.Itoa(banId))
  	utils.MemberCanRestrictMembers(b, u)
  	banMember, err := u.Message.Chat.GetMember(banId)
  	if banMember != nil {
    	if utils.MemberIsAdministrator(banMember) {
      	_, err = b.SendMessage(u.Message.Chat.Id, "У администраторов всегда карма чистая. По крайней мере здесь.")
      	return err
    	}
  	}
  	err = sqlite.ResetUserWarns(u.Message.Chat.Id, banId)
  	if err != nil {
    return err
  }
  	_, err = b.SendMessage(u.Message.Chat.Id, "У пользователя " + banMember.User.FirstName + " очищена карма.")
  	return err
}
