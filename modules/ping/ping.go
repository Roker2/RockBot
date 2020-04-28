package ping

import (
	"context"
	"fmt"
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/glinton/ping"
)

func Ping(b ext.Bot, u *gotgbot.Update, args []string) error {
	var msgText string
	if len(args) > 0 {
		for _, url := range args {
			res, err := ping.IPv4(context.Background(), url)
			if err != nil {
				return err
			}
			msgText += fmt.Sprintf("Completed one ping to %s with %d bytes in %v\n", url,
				res.TotalLength, res.RTT)
		}
	} else {
		res, err := ping.IPv4(context.Background(), "google.com")
		if err != nil {
			return err
		}
		msgText = fmt.Sprintf("Completed one ping to google.com with %d bytes in %v\n",
			res.TotalLength, res.RTT)
	}
	_, err := b.SendMessage(u.Message.Chat.Id, msgText)
	return err
}
