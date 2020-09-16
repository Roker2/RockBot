package welcome

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"github.com/Roker2/RockBot/modules/errors"
	"github.com/Roker2/RockBot/modules/sql"
	"github.com/Roker2/RockBot/modules/texts"
	"github.com/Roker2/RockBot/modules/utils"
	"strings"
)

//It add some info about user to text
func textHandler(text string, user *ext.User) string {
	text = strings.ReplaceAll(text, "<br>", "\n")
	//Replace {firstName} to first name of user
	if user.FirstName != "" {
		text = strings.ReplaceAll(text, "{firstName}", user.FirstName)
	} else {
		text = strings.ReplaceAll(text, "{firstName}", texts.User)
	}
	//Replace {lastName} to last name of user
	if user.LastName != "" {
		text = strings.ReplaceAll(text, "{lastName}", user.LastName)
	}
	//Replace {username} to last name of user
	if user.LastName != "" {
		text = strings.ReplaceAll(text, "{username}", user.Username)
	}
	return text
}

func NewMember(b ext.Bot, u *gotgbot.Update) error {
	newMembers := u.EffectiveMessage.NewChatMembers
	var prevUser *ext.User
	for _, member := range newMembers {
		if (member.Id == b.Id) || (prevUser != nil && prevUser.Id == member.Id) {
			continue
		}
		err := Welcome(b, u)
		if err != nil {
			return err
		}
		prevUser = &member
	}
	return nil
}

func LeftMember(b ext.Bot, u *gotgbot.Update) error {
	member := u.EffectiveMessage.LeftChatMember
	text := textHandler(texts.ByeUser, member)
	_, err := b.SendMessage(u.Message.Chat.Id, text)
	if err != nil {
		return err
	}
	return nil
}

func SetWelcome(b ext.Bot, u *gotgbot.Update, args []string) error {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	if len(args) == 0 {
		_, err := b.SendMessage(u.Message.Chat.Id, texts.AboutSetWelcome)
		return err
	}
	welcome := utils.RemoveCommand(u.Message.OriginalHTML())
	err = sql.SetWelcome(u.Message.Chat.Id, welcome)
	if err != nil {
		return err
	}
	_, err = b.SendMessageHTML(u.Message.Chat.Id, texts.NewWelcomeIsSettled)
	if err != nil {
		return err
	}
	err = Welcome(b, u)
	return err
}

func ResetWelcome(b ext.Bot, u *gotgbot.Update) error {
	member, err := u.Message.Chat.GetMember(u.Message.From.Id)
	if !utils.MemberIsAdministrator(member) {
		_, err = b.SendMessage(u.Message.Chat.Id, texts.YouAreNotAdministrator)
		return err
	}
	err = sql.SetWelcome(u.Message.Chat.Id, texts.DefaultWelcome)
	if err != nil {
		return err
	}
	_, err = b.SendMessageHTML(u.Message.Chat.Id, texts.NewWelcomeIsSettled)
	if err != nil {
		return err
	}
	err = Welcome(b, u)
	return err
}

func Welcome(b ext.Bot, u *gotgbot.Update) error {
	disabledCommands, err := sql.GetDisabledCommands(u.Message.Chat.Id)
	if err != nil {
		return err
	}
	if strings.Contains(disabledCommands, "welcome") {
		return nil
	}
	member := u.Message.From
	welcome, err := sql.GetWelcome(u.Message.Chat.Id)
	if err != nil {
		errors.SendError(err)
	}
	newMsg := b.NewSendableMessage(u.Message.Chat.Id, "")
	index := strings.Index(welcome, "[buttons]")
	if index != -1 {
		buttonsText := welcome[index + 10:]
		welcome = welcome[:index]
		//create empty markup
		markup := ext.InlineKeyboardMarkup{
			InlineKeyboard: &[][]ext.InlineKeyboardButton{},
		}
		//create two-dimensional array of buttons
		var inlineKeyboard [][]ext.InlineKeyboardButton
		//split \n
		//\n - new line of buttons
		newLineSplit := strings.Split(buttonsText, "\n")
		for _, temp1 := range newLineSplit {
			//create array of buttons (one line)
			var tempMassive []ext.InlineKeyboardButton
			//split ", "
			//it is separation to buttons from line
			commaSplit := strings.Split(temp1, ", ")
			for _, temp2 := range commaSplit {
				splittedText := strings.Split(temp2, " - ")
				//add button to line
				tempMassive = append(tempMassive, ext.InlineKeyboardButton{Text:splittedText[0], Url:splittedText[1]})
			}
			//add line to array of buttons
			inlineKeyboard = append(inlineKeyboard, tempMassive)
		}
		//set array of buttons
		markup.InlineKeyboard = &inlineKeyboard
		newMsg.ReplyMarkup = ext.ReplyMarkup(&markup)
	}
	newMsg.Text = textHandler(welcome, member)
	newMsg.ParseMode = parsemode.Html
	_, err = newMsg.Send()
	return err
}