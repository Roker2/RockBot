package main

import (
	"./modules/admin"
	"./modules/bans"
	"./modules/info"
	"./modules/media/Anilibria"
	"./modules/mute"
	"./modules/welcome"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"log"
	"fmt"
	"encoding/json"
	"os"
)

type Config struct {
    TelegramBotToken string
}

func start(b ext.Bot, u *gotgbot.Update) error {
	_, err := b.SendMessage(u.Message.Chat.Id, "Привет. Меня зовут Рок Драгоций, я являюсь чат-ботом и книжным персонажем. Если что-то понадобится, то просто введите команду.")
	return err
}

/*func test(b ext.Bot, u *gotgbot.Update) error {
	_, err := b.SendMessage(u.Message.Chat.Id, strconv.FormatBool(utils.MemberCanRestrictMembers(b, u)))
	return err
}*/

func main() {
	log.Println("Starting Rock...")
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(configuration.TelegramBotToken)
	updater, err := gotgbot.NewUpdater(configuration.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}
	updater.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	updater.Dispatcher.AddHandler(handlers.NewCommand("randomal", Anilibria.Randomal))
	updater.Dispatcher.AddHandler(handlers.NewCommand("info", info.UserInfo))
	updater.Dispatcher.AddHandler(handlers.NewCommand("chatinfo", info.ChatInfo))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("ban", []rune{'/', '!'}, bans.Ban))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("unban", []rune{'/', '!'}, bans.Unban))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("kick", []rune{'/', '!'}, bans.Kick))
	updater.Dispatcher.AddHandler(handlers.NewCommand("kickme", bans.Kickme))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("mute", []rune{'/', '!'}, mute.Mute))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("unmute", []rune{'/', '!'}, mute.Unmute))
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.NewChatMembers(), welcome.NewMember))
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.LeftChatMembers(), welcome.LeftMember))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("pin", []rune{'/', '!'}, admin.Pin))
	updater.Dispatcher.AddHandler(handlers.NewCommand("unpin", admin.Unpin))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("promote", []rune{'/', '!'}, admin.Promote))
	updater.Dispatcher.AddHandler(handlers.NewPrefixArgsCommand("demote", []rune{'/', '!'}, admin.Demote))
	updater.Dispatcher.AddHandler(handlers.NewCommand("purge", admin.Purge))
	//updater.Dispatcher.AddHandler(handlers.NewCommand("test", test))
	log.Println("Starting long polling")
	err = updater.StartPolling()
	if err != nil {
		log.Fatal(err)
	}
	updater.Idle()
	log.Println("Rock is started.")
}
