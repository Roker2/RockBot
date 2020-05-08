package warns
//version: 1,0

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/Roker2/RockBot/modules/bans"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
	"log"
	"strconv"
)

func WarnUser(b ext.Bot, u *gotgbot.Update, args []string) error {
	canBan, banId, err := utils.CommonBan(b, u, args)
	if !canBan {
		return err
	}
 	banMember, err := u.Message.Chat.GetMember(banId)
 	if err != nil {
 		return err
	}
 	if utils.MemberIsAdministrator(banMember) {
 		_, err = b.SendMessage(u.Message.Chat.Id, "Я не могу дать предупреждение администратору.")
 		return err
 	}
	maxQuantity, err := sql.GetWarnsQuantityOfChat(u.Message.Chat.Id)
	if err != nil {
		return err
	}
  	quantity, err := sql.AddUserWarn(u.Message.Chat.Id, banId)
  	if err != nil {
    	return err
  	}
  	_, err = b.SendMessage(u.Message.Chat.Id, "Количество предупреждений у " + banMember.User.FirstName + ": " + strconv.Itoa(quantity) + "/" + strconv.Itoa(maxQuantity))
  	if err != nil {
  		return err
  	}
  	if quantity >= maxQuantity {
  		err = bans.Ban(b, u, args)
  		return err
  	}
  	return nil
}

/*func WarnsQuantity (b ext.Bot, u *gotgbot.Update) error {
  	_, err := b.SendMessage(u.Message.Chat.Id, strconv.Itoa(sql.GetWarnsQuantityOfChat(u.Message.Chat.Id)))
  	return err
}*/

func GetUserWarns(b ext.Bot, u *gotgbot.Update, args []string) error {
	banId, _ := utils.ExtractId(b, u, args)
	if banId == 0 {
		banId = u.Message.From.Id
	}
 	if utils.ItIsMe(b, u, banId) {
 		return nil
 	}
 	log.Print(strconv.Itoa(banId))
	UserWarns, err := sql.GetUserWarns(u.Message.Chat.Id, banId)
	if err != nil {
		return err
  	} else {
 	  	banMember, err := u.Message.Chat.GetMember(banId)
    	if err != nil {
      		return err
    	}
		maxQuantity, err := sql.GetWarnsQuantityOfChat(u.Message.Chat.Id)
		if err != nil {
			return err
		}
    	_, err = b.SendMessage(u.Message.Chat.Id, "Количество предупреждений у " + banMember.User.FirstName + ": " + strconv.Itoa(UserWarns) + "/" + strconv.Itoa(maxQuantity))
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
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
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
		err = sql.SetWarnsQuantityOfChat(u.Message.Chat.Id, quantity)
		if err != nil {
			return err
		}
		_, err = b.SendMessage(u.Message.Chat.Id, "Новое количество максимальных предупреждений: " + args[0] + ".")
		if err != nil {
			return err
		}
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
    	_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
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
  	err = sql.ResetUserWarns(u.Message.Chat.Id, banId)
  	if err != nil {
    return err
  }
  	_, err = b.SendMessage(u.Message.Chat.Id, "У пользователя " + banMember.User.FirstName + " очищена карма.")
  	return err
}
