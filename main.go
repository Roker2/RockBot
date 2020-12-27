package main

import (
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"github.com/Roker2/RockBot/modules/admin"
	"github.com/Roker2/RockBot/modules/bans"
	"github.com/Roker2/RockBot/modules/info"
	"github.com/Roker2/RockBot/modules/media/Anilibria"
	"github.com/Roker2/RockBot/modules/media/SunMyungMoon"
	"github.com/Roker2/RockBot/modules/mute"
	"github.com/Roker2/RockBot/modules/ping"
	"github.com/Roker2/RockBot/modules/rules"
	"github.com/Roker2/RockBot/modules/warns"
	"github.com/Roker2/RockBot/modules/welcome"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

func start(b ext.Bot, u *gotgbot.Update) error {
	_, err := b.SendMessage(u.Message.Chat.Id, "Привет. Меня зовут Рок Драгоций, я являюсь чат-ботом и книжным персонажем. Если что-то понадобится, то просто введите команду.")
	return err
}

/*func test(b ext.Bot, u *gotgbot.Update) error {
	_, err := b.SendMessage(u.Message.Chat.Id, strconv.FormatBool(utils.MemberCanRestrictMembers(b, u)))
	return err
}*/

func main() {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))
	defer logger.Sync() // flushes buffer, if any
	l := logger.Sugar()

	l.Info("Starting Rock...")

	l.Info(os.Getenv("TOKEN"))
	updater, err := gotgbot.NewUpdater(logger, os.Getenv("TOKEN"))
	if err != nil {
		l.Fatalw("failed to start updater", zap.Error(err))
	}
	updater.Dispatcher.AddHandler(handlers.NewCallback("removeWarn", warns.RemoveWarnButton))
	updater.Dispatcher.AddHandler(handlers.NewCommand("start", start))
	updater.Dispatcher.AddHandler(handlers.NewCommand("randomal", Anilibria.Randomal))
	updater.Dispatcher.AddHandler(handlers.NewCommand("randomsmmq", SunMyungMoon.RandomSMMQ))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("info", info.UserInfo))
	updater.Dispatcher.AddHandler(handlers.NewCommand("chatinfo", info.ChatInfo))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("ban", bans.Ban))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("unban", bans.Unban))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("kick", bans.Kick))
	updater.Dispatcher.AddHandler(handlers.NewCommand("kickme", bans.Kickme))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("mute", mute.Mute))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("tmute", mute.TemporaryMute))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("unmute", mute.Unmute))
	updater.Dispatcher.AddHandler(handlers.NewCommand("welcome", welcome.Welcome))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("setwelcome", welcome.SetWelcome))
	updater.Dispatcher.AddHandler(handlers.NewCommand("resetwelcome", welcome.ResetWelcome))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("pin", admin.Pin))
	updater.Dispatcher.AddHandler(handlers.NewCommand("unpin", admin.Unpin))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("promote", admin.Promote))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("demote", admin.Demote))
	updater.Dispatcher.AddHandler(handlers.NewCommand("purge", admin.Purge))
	updater.Dispatcher.AddHandler(handlers.NewCommand("report", admin.Report))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("warn", warns.WarnUser))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("warns", warns.GetUserWarns))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("setwarns", warns.SetWarnsQuantity))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("resetwarns", warns.ResetWarns))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("setrules", rules.SetRules))
	updater.Dispatcher.AddHandler(handlers.NewCommand("rules", rules.GetRules))
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("ping", ping.Ping))//heroku doesn't support ping :(
	updater.Dispatcher.AddHandler(handlers.NewArgsCommand("disablecommands", admin.DisableCommands))
	updater.Dispatcher.AddHandler(handlers.NewCommand("disabledcommands", admin.GetDisabledCommands))
	updater.Dispatcher.AddHandler(handlers.NewCommand("disableallcommands", admin.DisableAllCommands))
	updater.Dispatcher.AddHandler(handlers.NewCommand("enableallcommands", admin.EnableAllCommands))
	//updater.Dispatcher.AddHandler(handlers.NewCommand("test", test))
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.NewChatMembers(), welcome.NewMember))
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.LeftChatMembers(), welcome.LeftMember))
	updater.Dispatcher.AddHandler(handlers.NewMessage(Filters.All, info.SaveUserToDatabase))
	// start getting updates
	if os.Getenv("USE_WEBHOOKS") == "yes" {
		logrus.Println("Starting webhook")
		port, err := strconv.Atoi(os.Getenv("PORT"))
		herokuUrl := os.Getenv("HEROKU_URL")
		webhook := ext.Webhook{
			Serve:          "0.0.0.0",
			ServePort:      port,
			ServePath:      updater.Bot.Token,
			URL:            herokuUrl,
			MaxConnections: 40,
		}
		updater.StartWebhook(webhook)
		ok, err := updater.SetWebhook(updater.Bot.Token, webhook)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to start bot due to: ", err)
		}
		if !ok {
			logrus.Fatal("Failed to set webhook")
		}
	} else {
		logrus.Println("Starting long polling")
		err = updater.StartPolling()
		if err != nil {
			logrus.Fatal(err)
		}
	}
	updater.Idle()
	logrus.Println("Rock is started.")
}
