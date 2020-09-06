package Anilibria

import (
	"encoding/json"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"github.com/Roker2/RockBot/modules/sql"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ALRelease struct {
	code string
	name string
	series string
	releaseType string
	genres []string
	description string
	pictureurl string
}

func (release ALRelease) ToString() string {
	var str string
	str = "<b>Название:</b> " + release.name
	if release.series != "1" {
		str += "\n<b>Серии:</b> " + release.series
	}
	str += "\n<b>Тип релиза:</b> " + release.releaseType
	str += "\n<b>Жанры:</b> "
	for _, value := range release.genres {
		str += value + ", "
	}
	str = strings.TrimSuffix(str, " ")
	str = strings.TrimSuffix(str, ",")
	str += "\n<b>Описание:</b> " + release.description
	//str += "\n<a href=\"https://www.anilibria.tv/release/" + release.code + ".html\">Ссылка</a>"
	return str
}

func toALRelease(info map[string]interface{}) ALRelease {
	var AlRel ALRelease
	AlRel.code = info["code"].(string)
	AlRel.name = info["names"].([]interface{})[0].(string)
	AlRel.series = info["series"].(string)
	AlRel.releaseType = info["type"].(string)
	genres := info["genres"].([]interface{})
	for _, value := range genres {
		AlRel.genres = append(AlRel.genres, value.(string))
	}
	AlRel.description = info["description"].(string)
	AlRel.pictureurl = info["poster"].(string)
	return  AlRel
}

func Randomal(b ext.Bot, u *gotgbot.Update) error  {
	disabledCommands, err := sql.GetDisabledCommands(u.Message.Chat.Id)
	if err != nil {
		return err
	}
	if strings.Contains(disabledCommands, "randomal") {
		return nil
	}
	resp, err := http.PostForm("https://www.anilibria.tv/public/api/index.php",
		url.Values{"query": {"random_release"}})
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var randomRel map[string]interface{}
	err = json.Unmarshal([]byte(body), &randomRel)
	if err != nil {
		return err
	}
	resp, err = http.PostForm("https://www.anilibria.tv/public/api/index.php",
		url.Values{"query": {"release"}, "code": {(randomRel["data"].(map[string]interface{}))["code"].(string)}})
	if err != nil {
		return err
	}
	if resp.Body == nil {
		log.Printf("resp is nil")
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(body), &randomRel)
	if err != nil {
		return err
	}
	release := toALRelease(randomRel["data"].(map[string]interface{}))
	release.description = strings.Replace(release.description, "<br>", "\n", -1)
	//log.Print(release.ToString())
	//str = str[:1000 - len("https://www.anilibria.tv/release/" + release.code + ".html") - 25]
	markup := ext.InlineKeyboardMarkup{
		InlineKeyboard: &[][]ext.InlineKeyboardButton{
			[]ext.InlineKeyboardButton{
				ext.InlineKeyboardButton{Text:"Ссылка", Url:"https://www.anilibria.tv/release/" + release.code + ".html"},
			},
		},
	}
	if len(release.ToString()) <= 1024 {
		msg := b.NewSendablePhoto(u.Message.Chat.Id, release.ToString())
		msg.ParseMode = parsemode.Html
		msg.Photo = b.NewFileId("https://www.anilibria.tv" + release.pictureurl)
		msg.ReplyMarkup = ext.ReplyMarkup(&markup)
		_, err = msg.Send()
		if err != nil {
			return err
		}
	} else {
		msg := b.NewSendablePhoto(u.Message.Chat.Id, "")
		msg.Photo = b.NewFileId("https://www.anilibria.tv" + release.pictureurl)
		_, err = msg.Send()
		if err != nil {
			return err
		}
		msgtext := b.NewSendableMessage(u.Message.Chat.Id, release.ToString())
		msgtext.ParseMode = parsemode.Html
		msgtext.ReplyMarkup = ext.ReplyMarkup(&markup)
		_, err = msgtext.Send()
		if err != nil {
			return err
		}
	}
	return nil
}
